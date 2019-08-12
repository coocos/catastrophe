package main

import (
	"github.com/coocos/catastrophe/feed"
	"log"
)

func main() {

	rawFeed, err := feed.FetchFeed()
	if err != nil {
		log.Fatalln("Failed to retrieve event feed")
	}

	items := feed.Parse(rawFeed)
	for _, item := range items {
		log.Printf("%s", item)
	}
}
