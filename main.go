package main

import (
	"github.com/coocos/catastrophe/feed"
	"log"
	"time"
)

func Update(ticker *time.Ticker) {

	currentEvent := feed.Event{}

	for _ = range ticker.C {
		latestEvents, err := feed.LatestEvents()
		if err != nil {
			log.Printf("Failed to retrieve event feed: %s", err)
			continue
		}
		latestEvent := latestEvents[len(latestEvents)-1]

		// TODO: This might swallow events if there were multiple new events
		if latestEvent.Timestamp != currentEvent.Timestamp {
			log.Printf("%s", latestEvent)
			currentEvent = latestEvent
		}
	}

}

func main() {

	ticker := time.NewTicker(60 * 1000 * time.Millisecond)

	go Update(ticker)

	time.Sleep(60 * 10000 * time.Millisecond)
	ticker.Stop()

}
