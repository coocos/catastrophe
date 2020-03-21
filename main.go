package main

import (
	"net/http"
	"sync"
	"time"

	"github.com/coocos/catastrophe/feed"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

var eventStream = make(chan *feed.Event)
var upgrader = websocket.Upgrader{}
var clients = make(map[*websocket.Conn]bool)
var clientMutex = &sync.Mutex{}
var latestEvent = feed.Event{}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Warn("Failed to upgrade WebSocket connection")
		return
	}
	log.Info("New client connected")

	clientMutex.Lock()
	clients[conn] = true
	clientMutex.Unlock()

	conn.WriteJSON(latestEvent)
}

func UpdateEvents(ticker *time.Ticker, eventStream chan<- *feed.Event) {

	client := feed.NewClient()

	allEvents, err := client.LatestEvents()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("Failed to retrieve starting event")
	} else {
		latestEvent = allEvents[len(allEvents)-1]
		eventStream <- &latestEvent
	}

	for _ = range ticker.C {
		newEvents, err := client.EventsSince(latestEvent.Time)
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("Failed to retrieve event feed")
			continue
		}

		log.WithFields(log.Fields{
			"events": len(newEvents),
		}).Info("Found new events")
		for _, event := range newEvents {
			eventStream <- &event
		}
		if len(newEvents) > 0 {
			latestEvent = newEvents[len(newEvents)-1]
		}
	}

}

func BroadcastEvents(eventStream <-chan *feed.Event) {

	for event := range eventStream {
		clientMutex.Lock()
		log.WithFields(log.Fields{
			"connected_clients": len(clients),
		}).Info("Broadcasting event to clients")
		log.Info(event)
		for client := range clients {
			err := client.WriteJSON(event)
			if err != nil {
				log.Info("Dropping client")
				delete(clients, client)
			}
		}
		clientMutex.Unlock()
	}
}

func configureLogger() {
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

func main() {

	configureLogger()

	ticker := time.NewTicker(60 * 1000 * time.Millisecond)
	go UpdateEvents(ticker, eventStream)
	go BroadcastEvents(eventStream)

	http.HandleFunc("/websocket", WebSocketHandler)
	http.Handle("/", http.FileServer(http.Dir("./frontend/dist")))
	http.ListenAndServe("localhost:8080", nil)

}
