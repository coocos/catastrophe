package server

import (
	"reflect"
	"testing"
	"time"

	"github.com/coocos/catastrophe/internal/feed"
)

type TestConnection struct {
	event  *feed.Event
	closed bool
}

func (c *TestConnection) WriteJSON(v interface{}) error {
	event, ok := v.(*feed.Event)
	if ok {
		c.event = event
	}
	return nil
}

func (c *TestConnection) Close() error {
	c.closed = true
	return nil
}

func createGroupWithConnection() (*ConnectionGroup, *TestConnection) {
	group := NewConnectionGroup()
	connection := &TestConnection{}
	group.Add(connection)
	return group, connection
}

// TestConnectionGroup tests groups of connections
func TestConnectionGroup(t *testing.T) {

	t.Run("test adding connection to group", func(t *testing.T) {
		group, _ := createGroupWithConnection()
		if group.Count() != 1 {
			t.Errorf("Group should have 1 connection, it has %d", group.Count())
		}
	})

	t.Run("test closing down group", func(t *testing.T) {
		group, connection := createGroupWithConnection()
		group.Shutdown()
		if !connection.closed {
			t.Errorf("Connection should have been closed")
		}
	})

	t.Run("test broadcasting to group", func(t *testing.T) {
		group, connection := createGroupWithConnection()
		event := &feed.Event{
			Type:     "Tulipalo",
			Location: "Helsinki",
			Time:     time.Now().UTC(),
		}
		group.Broadcast(event)
		if !reflect.DeepEqual(connection.event, event) {
			t.Errorf("Broadcasted %v, expected %v", connection.event, event)
		}
	})

}
