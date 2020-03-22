package main

import (
	"flag"
	"time"

	"github.com/coocos/catastrophe/feed"
	"github.com/coocos/catastrophe/server"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

var latestEvent = feed.Event{}

func pollEvents(client *feed.Client, ticker *time.Ticker, eventStream chan<- *feed.Event) {

	events, err := client.LatestEvents()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Fatal("Failed to retrieve starting event")
	}

	latestEvent = events[len(events)-1]
	eventStream <- &latestEvent

	for _ = range ticker.C {
		newEvents, err := client.EventsSince(latestEvent.Time)
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("Failed to retrieve event feed")
			continue
		}
		if len(newEvents) == 0 {
			continue
		}
		latestEvent = newEvents[len(newEvents)-1]

		log.WithFields(log.Fields{
			"events": len(newEvents),
		}).Info("Found new events")
		for _, event := range newEvents {
			eventStream <- &event
		}
	}

}

func broadcastEvents(eventStream <-chan *feed.Event, server *server.WebSocketServer) {

	for event := range eventStream {
		log.WithFields(log.Fields{
			"event_location": event.Location,
			"event_type":     event.Type,
			"event_time":     event.Time,
		}).Info("Publishing new event")
		server.Send(event)
	}

}

func configureLogger() {
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

func main() {

	configureLogger()

	host := flag.String("host", "localhost", "host to use")
	port := flag.Int("port", 8000, "port to use")
	flag.Parse()

	server := server.NewServer(*host, *port)
	defer server.Start()

	eventStream := make(chan *feed.Event)
	go pollEvents(feed.NewClient(), time.NewTicker(60*time.Second), eventStream)
	go broadcastEvents(eventStream, server)

}
