package server

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/coocos/catastrophe/feed"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

// WebSocketServer is used to interface with WebSocket connections
type WebSocketServer struct {
	Clients      map[*websocket.Conn]bool
	clientMutex  *sync.Mutex
	initialEvent *feed.Event
	Port         int
	Host         string
}

// Send sends event to all connected WebSocket clients
func (server WebSocketServer) Send(event *feed.Event) {

	server.initialEvent = event

	server.clientMutex.Lock()
	defer server.clientMutex.Unlock()

	for client := range server.Clients {
		err := client.WriteJSON(event)
		if err != nil {
			log.Info("Dropping client")
			delete(server.Clients, client)
		}
		log.WithFields(log.Fields{
			"clients": len(server.Clients),
			"event":   *event,
		}).Info("Sending event to clients")
	}

}

// NewServer returns a new WebSocketServer
func NewServer(host string, port int) *WebSocketServer {
	server := WebSocketServer{
		Clients:     make(map[*websocket.Conn]bool),
		clientMutex: &sync.Mutex{},
		Port:        port,
		Host:        host,
	}
	return &server
}

// Start starts the WebSocket server and blocks
func (server WebSocketServer) Start() {

	upgrader := websocket.Upgrader{}
	http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Warn("Failed to upgrade WebSocket connection")
			return
		}

		server.clientMutex.Lock()
		defer server.clientMutex.Unlock()
		log.WithFields(log.Fields{
			"connected_clients": len(server.Clients),
		}).Info("New client connected")
		server.Clients[conn] = true

		conn.WriteJSON(server.initialEvent)
	})

	log.WithFields(log.Fields{
		"host": server.Host,
		"port": server.Port,
	}).Info("Starting server")
	http.ListenAndServe(fmt.Sprintf("%s:%d", server.Host, server.Port), nil)

}
