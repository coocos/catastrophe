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
	httpServer  *http.Server
	lock        *sync.Mutex
	latestEvent *feed.Event
	upgrader    *websocket.Upgrader
	port        int
	host        string
}

// NewWebSocketServer constructs a WebSocketServer
func NewWebSocketServer(host string, port int) *WebSocketServer {
	server := WebSocketServer{
		connections: make(map[*websocket.Conn]bool),
		lock:        &sync.Mutex{},
		port:        port,
		host:        host,
		upgrader:    &websocket.Upgrader{},
	}

	mux := http.NewServeMux()
	mux.Handle("/websocket", http.HandlerFunc(server.handleNewConnection))

	server.httpServer = &http.Server{
		Addr: fmt.Sprintf("%s:%d", host, port),
		Handler: mux,
	}

	return &server
}

// Publish sends event to all connected WebSocket clients
func (s *WebSocketServer) Publish(event *feed.Event) {

	s.lock.Lock()
	defer s.lock.Unlock()

	log.WithFields(log.Fields{
		"event_location": event.Location,
		"event_type":     event.Type,
		"event_time":     event.Time,
		"connections":    len(s.connections),
	}).Info("Publishing new event")

	s.latestEvent = event

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

	s.lock.Lock()
	defer s.lock.Unlock()
	s.connections[conn] = true
	log.WithFields(log.Fields{
		"connections": len(s.connections),
	}).Info("New connection")

}

func (s *WebSocketServer) handleNewConnection(w http.ResponseWriter, r *http.Request) {

	conn, err := s.upgrader.Upgrade(w, r, nil)
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

	s.lock.Lock()
	for connection := range s.connections {
		connection.Close()
	}
	s.lock.Unlock()

	s.httpServer.Close()
}

// Start starts the WebSocket server and blocks
func (s *WebSocketServer) Start() {

	log.WithFields(log.Fields{
		"host": s.host,
		"port": s.port,
	}).Info("Starting server")

	if err := s.httpServer.ListenAndServe(); err != http.ErrServerClosed {
		log.WithFields(log.Fields{
			"host": s.host,
			"port": s.port,
		}).Fatal("Failed to start server - is the port in use?")
	}

}
