package feed

import (
	"time"

	log "github.com/sirupsen/logrus"
)

// PollEvents polls the event feed and dispatches the events to the server for publishing
func PollEvents(client Client, ticker *time.Ticker, eventStream chan<- *Event) {

	events, err := client.LatestEvents()
	// Fail fast at startup
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Fatal("Failed to retrieve starting event")
	}

	latestEvent := events[len(events)-1]
	eventStream <- &latestEvent

	// Periodically check the feed for new events
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
		log.WithFields(log.Fields{
			"events": len(newEvents),
		}).Info("Found new events")

		for _, event := range newEvents {
			eventStream <- &event
		}
		latestEvent = newEvents[len(newEvents)-1]
	}

}
