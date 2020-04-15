package feed

import (
	"time"

	log "github.com/sirupsen/logrus"
)

func getFirstEvent(client Client) Event {
	events, err := client.LatestEvents()
	// Fail fast if unable to fetch the first event
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Fatal("Failed to retrieve starting event")
	}
	return events[len(events)-1]
}

// PollEvents polls the event feed and dispatches the events to the server for publishing
func PollEvents(client Client, ticker *time.Ticker, eventStream chan<- *Event) {

	latestEvent := getFirstEvent(client)
	eventStream <- &latestEvent

	// Periodically check the feed for new events
	// TODO: Maybe this function should not even be aware of the ticker... Instead it
	// use select with two channels, one that yields commands to grab the feed and the
	// other channel to quit?
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
