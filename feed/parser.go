package feed

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"strings"
)

type Event struct {
	Type        string
	Location    string
	Timestamp   string //TODO: Replace with a date type
	Description string
}

func Parse(rawFeed string) []Event {
	parser := gofeed.NewParser()
	feed, err := parser.ParseString(rawFeed)
	if err != nil {
		fmt.Println(err)
	}

	events := []Event{}

	for _, item := range feed.Items {
		event := Event{
			Type:        strings.Split(item.Title, ", ")[1],
			Location:    strings.Split(item.Title, ", ")[0],
			Timestamp:   strings.Split(item.Description, " ")[1],
			Description: item.Description,
		}
		events = append(events, event)
	}

	return events
}
