package feed

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

const feedUrl string = "http://www.peto-media.fi/tiedotteet/rss.xml"

func FetchFeed() (string, error) {

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
