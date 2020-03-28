package server

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/coocos/catastrophe/feed"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

// EventServer is an interface for servers which can publish rescue service events
type EventServer interface {
	Publish(*feed.Event)
	Start()
}

// WebSocketServer is used to interface with WebSocket connections
type WebSocketServer struct {
	Clients     map[*websocket.Conn]bool
	clientMutex *sync.Mutex
	latestEvent *feed.Event
	Port        int
	Host        string
}

// NewWebSocketServer constructs a WebSocketServer
func NewWebSocketServer(host string, port int) *WebSocketServer {
	server := WebSocketServer{
		Clients:     make(map[*websocket.Conn]bool),
		clientMutex: &sync.Mutex{},
		Port:        port,
		Host:        host,
	}
	return &server
}

// Publish sends event to all connected WebSocket clients
func (s WebSocketServer) Publish(event *feed.Event) {

	s.latestEvent = event

	s.clientMutex.Lock()
	defer s.clientMutex.Unlock()

	log.WithFields(log.Fields{
		"event_location": event.Location,
		"event_type":     event.Type,
		"event_time":     event.Time,
		"clients":        len(s.Clients),
	}).Info("Publishing new event")

	for client := range s.Clients {
		err := client.WriteJSON(event)
		if err != nil {
			log.Info("Dropping client")
			delete(s.Clients, client)
		}
	}

}

// Start starts the WebSocket server and blocks
func (s WebSocketServer) Start() {

	upgrader := websocket.Upgrader{}
	http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Warn("Failed to upgrade WebSocket connection")
			return
		}

		s.clientMutex.Lock()
		defer s.clientMutex.Unlock()
		log.WithFields(log.Fields{
			"connected_clients": len(s.Clients),
		}).Info("New client connected")
		s.Clients[conn] = true

		conn.WriteJSON(s.latestEvent)
	})

	log.WithFields(log.Fields{
		"host": s.Host,
		"port": s.Port,
	}).Info("Starting server")
	http.ListenAndServe(fmt.Sprintf("%s:%d", s.Host, s.Port), nil)

}
