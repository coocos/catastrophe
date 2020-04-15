package server

import (
	"sync"

	"github.com/coocos/catastrophe/internal/feed"
	log "github.com/sirupsen/logrus"
)

// Connection represents a single connection in the group
type Connection interface {
	WriteJSON(interface{}) error
	Close() error
}

// ConnectionGroup is a group of WebSocket connections
type ConnectionGroup struct {
	connections map[Connection]bool
	lock        *sync.Mutex
	cachedEvent *feed.Event
}

// NewConnectionGroup creates a new empty group
func NewConnectionGroup() *ConnectionGroup {
	return &ConnectionGroup{
		connections: make(map[Connection]bool),
		lock:        &sync.Mutex{},
	}
}

// Broadcast broadcasts event to all connections in the group
func (g *ConnectionGroup) Broadcast(event *feed.Event) {
	g.cachedEvent = event
	g.lock.Lock()
	defer g.lock.Unlock()
	for connection := range g.connections {
		err := connection.WriteJSON(event)
		if err != nil {
			log.Info("Dropping connection")
			delete(g.connections, connection)
			connection.Close()
		}
	}
}

// Count returns how many connections are currently in the group
func (g *ConnectionGroup) Count() int {
	g.lock.Lock()
	defer g.lock.Unlock()
	return len(g.connections)
}

// Shutdown shuts the group down by closing all connections
func (g *ConnectionGroup) Shutdown() {
	g.lock.Lock()
	for connection := range g.connections {
		connection.Close()
	}
	g.lock.Unlock()
}

// Add adds a connection to the group
func (g *ConnectionGroup) Add(connection Connection) {
	g.lock.Lock()
	defer g.lock.Unlock()
	g.connections[connection] = true
	if g.cachedEvent != nil {
		connection.WriteJSON(g.cachedEvent)
	}
}
