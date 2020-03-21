package feed

import (
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
	log "github.com/sirupsen/logrus"
)

type Event struct {
	Type        string    `json:"type"`
	Location    string    `json:"location"`
	Time        time.Time `json:"time"`
	Description string    `json:"description"`
}

func parseTimestamp(timestamp string) time.Time {

	date, err := time.Parse(time.RFC1123Z, timestamp)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("Failed to parse timestamp")
		date = time.Now()
	}
	return date

}

func Parse(rawFeed string) ([]Event, error) {
	parser := gofeed.NewParser()
	feed, err := parser.ParseString(rawFeed)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("Failed to parse events")
		return []Event{}, err
	}

	events := []Event{}

	//The latest events are first so reverse the order
	for i := len(feed.Items) - 1; i >= 0; i-- {
		item := feed.Items[i]
		event := Event{
			Type:        strings.Split(item.Title, ", ")[1],
			Location:    strings.Split(strings.Split(item.Title, ", ")[0], "/")[0],
			Time:        parseTimestamp(item.Published),
			Description: item.Description,
		}
		events = append(events, event)
	}

	return events, nil
}
