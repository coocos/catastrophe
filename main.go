package main

import (
	"github.com/coocos/catastrophe/feed"
	"log"
	"time"
)

func Update(ticker *time.Ticker) {

	allEvents, err := feed.LatestEvents()
	if err != nil {
		log.Printf("Failed to retrieve starting event %s", err)
	}
	latestEvent := allEvents[len(allEvents)-1]
	log.Printf("%s", latestEvent)

	for _ = range ticker.C {
		newEvents, err := feed.EventsSince(latestEvent.Timestamp)
		log.Printf("%d new event(s)", len(newEvents))
		if err != nil {
			log.Printf("Failed to retrieve event feed: %s", err)
			continue
		}
		for _, event := range newEvents {
			log.Printf("%s", event)
		}
	}

}

func main() {

	ticker := time.NewTicker(60 * 1000 * time.Millisecond)

	go Update(ticker)

	time.Sleep(60 * 10000 * time.Millisecond)
	ticker.Stop()

}
