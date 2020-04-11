package feed

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

const feedURL string = "http://www.peto-media.fi/tiedotteet/rss.xml"

// Client is an interface for clients which can be used to query events from a feed
type Client interface {
	LatestEvents() ([]Event, error)
	EventsSince(time.Time) ([]Event, error)
}

// RescueServiceClient is an HTTP client used to retrieve events from the Finnish rescue service API
type RescueServiceClient struct {
	baseURL string
}

// NewRescueServiceClient returns a pointer to a new Client instance
func NewRescueServiceClient() *RescueServiceClient {
	return &RescueServiceClient{
		feedURL,
	}
}

// LatestEvents returns a slice containing the latest published Events
func (c *RescueServiceClient) LatestEvents() ([]Event, error) {

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
func (c *RescueServiceClient) EventsSince(since time.Time) ([]Event, error) {

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

func (c *RescueServiceClient) fetchFeed() (string, error) {

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
