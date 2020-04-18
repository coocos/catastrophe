package feed

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

const feedURL string = "http://www.peto-media.fi/tiedotteet/rss.xml"

// Client is used to retrieve rescue service events
type Client interface {
	EventsSince(time.Time) ([]Event, error)
	LatestEvents() ([]Event, error)
}

// EventClient is a client used to retrieve rescue service events
type EventClient struct {
	baseURL string
}

// NewClient returns a new feed client
func NewClient() *EventClient {
	return &EventClient{
		feedURL,
	}
}

// LatestEvents returns a slice containing the latest published Events
func (c *EventClient) LatestEvents() ([]Event, error) {
	rawFeed, err := c.fetchFeed()
	if err != nil {
		return []Event{}, err
	}

	events, err := Parse(rawFeed)
	if err != nil {
		return []Event{}, err
	}

	return events, nil
}

// EventsSince returns a slice of Events that have occurred since the the passed time.Time
func (c *EventClient) EventsSince(since time.Time) ([]Event, error) {
	events, err := c.LatestEvents()
	if err != nil {
		return []Event{}, err
	}

	eventsSince := []Event{}
	for _, event := range events {
		if event.Time.After(since) {
			eventsSince = append(eventsSince, event)
		}
	}

	return eventsSince, nil
}

func (c *EventClient) fetchFeed() (string, error) {
	resp, err := http.Get(c.baseURL)
	if err != nil {
		log.Error("HTTP request to retrieve feed failed")
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.WithFields(log.Fields{
			"http_status": resp.StatusCode,
		}).Error("Failed to retrieve feed")
		return "", fmt.Errorf("Invalid HTTP response status code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Failed to read feed response")
		return "", err
	}

	return string(body), nil
}
