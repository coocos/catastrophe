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
	Shutdown()
}

// WebSocketServer is used to interface with WebSocket connections
type WebSocketServer struct {
	connections map[*websocket.Conn]bool
	server      *http.Server
	mutex       *sync.Mutex
	latestEvent *feed.Event
	port        int
	host        string
}

// NewWebSocketServer constructs a WebSocketServer
func NewWebSocketServer(host string, port int) *WebSocketServer {
	server := WebSocketServer{
		connections: make(map[*websocket.Conn]bool),
		mutex:       &sync.Mutex{},
		port:        port,
		host:        host,
	}
	return &server
}

// Publish sends event to all connected WebSocket clients
func (s *WebSocketServer) Publish(event *feed.Event) {

	s.latestEvent = event

	s.mutex.Lock()
	defer s.mutex.Unlock()

	log.WithFields(log.Fields{
		"event_location": event.Location,
		"event_type":     event.Type,
		"event_time":     event.Time,
		"connections":    len(s.connections),
	}).Info("Publishing new event")

	for connection := range s.connections {
		err := connection.WriteJSON(event)
		if err != nil {
			log.Info("Dropping connection")
			delete(s.connections, connection)
			connection.Close()
		}
	}

}

func (s *WebSocketServer) addConnection(conn *websocket.Conn) {

	s.mutex.Lock()
	defer s.mutex.Unlock()
	log.WithFields(log.Fields{
		"connections": len(s.connections),
	}).Info("New connection")
	s.connections[conn] = true

}

func (s *WebSocketServer) handleNewConnection(w http.ResponseWriter, r *http.Request) {

	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Warn("Failed to upgrade WebSocket connection")
		return
	}

	s.addConnection(conn)
	conn.WriteJSON(s.latestEvent)
}

// Shutdown closes all WebSocket connections and shuts down the server
func (s *WebSocketServer) Shutdown() {

	log.WithFields(log.Fields{
		"connections": len(s.connections),
	}).Info("Shutting down server")

	s.mutex.Lock()
	for connection := range s.connections {
		connection.Close()
	}
	s.mutex.Unlock()

	s.server.Close()
}

// Start starts the WebSocket server and blocks
func (s *WebSocketServer) Start() {

	log.WithFields(log.Fields{
		"host": s.host,
		"port": s.port,
	}).Info("Starting server")

	s.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.host, s.port),
		Handler: http.HandlerFunc(s.handleNewConnection),
	}

	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		log.WithFields(log.Fields{
			"host": s.host,
			"port": s.port,
		}).Fatal("Failed to start server - is the port in use?")
	}

}
