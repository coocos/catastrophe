package main

import (
	"github.com/coocos/catastrophe/feed"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"time"
)

var eventStream = make(chan *feed.Event)
var upgrader = websocket.Upgrader{}
var clients = make(map[*websocket.Conn]bool)
var clientMutex = &sync.Mutex{}
var latestEvent = feed.Event{}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade WebSocket connection")
		return
	}
	log.Printf("New client connected")

	clientMutex.Lock()
	clients[conn] = true
	clientMutex.Unlock()

	conn.WriteJSON(latestEvent)
}

func UpdateEvents(ticker *time.Ticker, eventStream chan<- *feed.Event) {

	client := feed.NewClient()

	allEvents, err := client.LatestEvents()
	if err != nil {
		log.Printf("Failed to retrieve starting event %s", err)
	}
	latestEvent = allEvents[len(allEvents)-1]
	eventStream <- &latestEvent

	for _ = range ticker.C {
		newEvents, err := client.EventsSince(latestEvent.Time)
		log.Printf("%d new event(s)", len(newEvents))
		if err != nil {
			log.Printf("Failed to retrieve event feed: %s", err)
			continue
		}
		for _, event := range newEvents {
			eventStream <- &event
		}
		if len(newEvents) > 0 {
			latestEvent = newEvents[len(newEvents)-1]
		}
	}

}

func BroadcastEvents(eventStream <-chan *feed.Event) {

	for event := range eventStream {
		log.Printf("%s", event)
		clientMutex.Lock()
		log.Printf("Broadcasting event to %d clients", len(clients))
		for client := range clients {
			err := client.WriteJSON(event)
			if err != nil {
				log.Printf("Dropping client")
				delete(clients, client)
			}
		}
		clientMutex.Unlock()
	}
}

func main() {

	ticker := time.NewTicker(60 * 1000 * time.Millisecond)
	go UpdateEvents(ticker, eventStream)
	go BroadcastEvents(eventStream)

	http.HandleFunc("/websocket", WebSocketHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe("localhost:8080", nil)

}
