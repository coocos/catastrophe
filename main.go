package main

import (
	"github.com/coocos/catastrophe/feed"
	"log"
	"time"
)

func Update(ticker *time.Ticker, eventStream chan<- *feed.Event) {

	client := feed.NewClient()

	allEvents, err := client.LatestEvents()
	if err != nil {
		log.Printf("Failed to retrieve starting event %s", err)
	}
	latestEvent := allEvents[len(allEvents)-1]
	eventStream <- &latestEvent

	for _ = range ticker.C {
		newEvents, err := client.EventsSince(latestEvent.Time)
		log.Printf("%d new event(s)", len(newEvents))
		if err != nil {
			log.Printf("Failed to retrieve event feed: %s", err)
			continue
		}
		for _, event := range newEvents {
			eventStream <- &event
		}
		if len(newEvents) > 0 {
			latestEvent = newEvents[len(newEvents)-1]
		}
	}

}

func main() {

	eventStream := make(chan *feed.Event)

	ticker := time.NewTicker(60 * 1000 * time.Millisecond)
	go Update(ticker, eventStream)

	for event := range eventStream {
		log.Printf("%s", event)
	}

}
