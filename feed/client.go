package feed

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const feedUrl string = "http://www.peto-media.fi/tiedotteet/rss.xml"

func LatestEvents() ([]Event, error) {

	rawFeed, err := fetchFeed()
	if err != nil {
		return []Event{}, err
	}

	events, err := Parse(rawFeed)
	if err != nil {
		return []Event{}, err
	}

	return events, nil

}

func EventsSince(since time.Time) ([]Event, error) {

	events, err := LatestEvents()
	if err != nil {
		return []Event{}, err
	}

	eventsSince := []Event{}
	for _, event := range events {
		if event.Timestamp.After(since) {
			eventsSince = append(eventsSince, event)
		}
	}

	return eventsSince, nil

}

func fetchFeed() (string, error) {

	resp, err := http.Get(feedUrl)
	if err != nil {
		log.Println("HTTP request to retrieve feed failed")
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Failed to retrieve feed: %s", resp.Status)
		return "", errors.New("Invalid HTTP response status code")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to read feed response")
		return "", err
	}

	return string(body), nil

}
