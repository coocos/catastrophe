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

// PollEvents returns a channel of feed.Events which yields new Events as they happen
func PollEvents(client Client, ticker *time.Ticker, stop <-chan bool) <-chan *Event {
	events := make(chan *Event)

	// Periodically check the feed for new events
	go func() {

		latestEvent := getFirstEvent(client)
		events <- &latestEvent

		for {
			select {
			case <-stop:
				close(events)
				return
			case <-ticker.C:
				newEvents, err := client.EventsSince(latestEvent.Time)
				if err != nil {
					log.WithFields(log.Fields{
						"err": err,
					}).Error("Failed to retrieve event feed")
					continue
				}
				if len(newEvents) > 0 {
					log.WithFields(log.Fields{
						"events": len(newEvents),
					}).Info("Found new events")
					for _, event := range newEvents {
						events <- &event
					}
					latestEvent = newEvents[len(newEvents)-1]
				}
			}

		}
	}()

	return events
}
