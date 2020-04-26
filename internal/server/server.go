package server

import (
	"fmt"
	"net/http"

	"github.com/coocos/catastrophe/internal/feed"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

// WebSocketServer is used to interface with WebSocket connections
type WebSocketServer struct {
	group      *ConnectionGroup
	httpServer *http.Server
	upgrader   *websocket.Upgrader
	port       int
	host       string
}

// NewWebSocketServer constructs a WebSocketServer
func NewWebSocketServer(host string, port int) *WebSocketServer {
	server := WebSocketServer{
		group:    NewConnectionGroup(),
		port:     port,
		host:     host,
		upgrader: &websocket.Upgrader{},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/websocket", server.handleNewConnection)

	server.httpServer = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: mux,
	}

	return &server
}

// Publish sends event to all connected WebSocket clients
func (s *WebSocketServer) Publish(event *feed.Event) {
	log.WithFields(log.Fields{
		"event_location": event.Location,
		"event_type":     event.Type,
		"event_time":     event.Time,
		"connections":    s.group.Count(),
	}).Info("Publishing new event")
	s.group.Broadcast(event)
}

func (s *WebSocketServer) handleNewConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Warn("Failed to upgrade WebSocket connection")
		return
	}
	s.group.Add(conn)
	log.WithFields(log.Fields{
		"connections": s.group.Count(),
	}).Info("New connection")
}

// Shutdown closes all WebSocket connections and shuts down the server
func (s *WebSocketServer) Shutdown() {
	log.WithFields(log.Fields{
		"connections": s.group.Count(),
	}).Info("Shutting down server")
	s.group.Shutdown()
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
