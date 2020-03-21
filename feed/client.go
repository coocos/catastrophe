package feed

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

const feedUrl string = "http://www.peto-media.fi/tiedotteet/rss.xml"

type Client struct {
	baseUrl string
}

func NewClient() *Client {
	return &Client{
		feedUrl,
	}
}

func (client *Client) LatestEvents() ([]Event, error) {

	rawFeed, err := client.fetchFeed()
	if err != nil {
		return []Event{}, err
	}

	events, err := Parse(rawFeed)
	if err != nil {
		return []Event{}, err
	}

	return events, nil

}

func (client *Client) EventsSince(since time.Time) ([]Event, error) {

	events, err := client.LatestEvents()
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

func (client *Client) fetchFeed() (string, error) {

	resp, err := http.Get(client.baseUrl)
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
