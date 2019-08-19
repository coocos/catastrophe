package feed

import (
	"github.com/mmcdole/gofeed"
	"log"
	"strings"
)

type Event struct {
	Type        string
	Location    string
	Timestamp   string //TODO: Replace with a date type
	Description string
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
			Timestamp:   strings.Split(item.Description, " ")[1],
			Description: item.Description,
		}
		events = append(events, event)
	}

	return events, nil
}
