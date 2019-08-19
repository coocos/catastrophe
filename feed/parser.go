package feed

import (
	"github.com/mmcdole/gofeed"
	"log"
	"strings"
	"time"
)

type Event struct {
	Type        string
	Location    string
	Timestamp   time.Time
	Description string
}

func parseTimestamp(timestamp string) time.Time {

	date, err := time.Parse(time.RFC1123Z, timestamp)
	if err != nil {
		log.Printf("Failed to parse timestamp: %s", err)
		date = time.Now()
	}
	return date

}

func Parse(rawFeed string) ([]Event, error) {
	parser := gofeed.NewParser()
	feed, err := parser.ParseString(rawFeed)
	if err != nil {
		log.Printf("Failed to parse events: %s", err)
		return []Event{}, err
	}

	events := []Event{}

	//The latest events are first so reverse the order
	for i := len(feed.Items) - 1; i >= 0; i-- {
		item := feed.Items[i]
		event := Event{
			Type:        strings.Split(item.Title, ", ")[1],
			Location:    strings.Split(strings.Split(item.Title, ", ")[0], "/")[0],
			Timestamp:   parseTimestamp(item.Published),
			Description: item.Description,
		}
		events = append(events, event)
	}

	return events, nil
}
