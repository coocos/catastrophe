package server

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/coocos/catastrophe/internal/feed"
	"github.com/gorilla/websocket"
)

func createEvent() *feed.Event {
	return &feed.Event{
		Type:     "Tulipalo",
		Location: "Helsinki",
		Time:     time.Now().UTC(),
	}
}

func createClient(host string, port int, t *testing.T) *websocket.Conn {
	url := url.URL{
		Scheme: "ws",
		Host:   fmt.Sprintf("%s:%d", host, port),
		Path:   "/websocket",
	}

	// Attempt to connect until the server is up and responding
	for {
		connection, response, err := websocket.DefaultDialer.Dial(url.String(), nil)
		if err != nil || response == nil {
			time.Sleep(10 * time.Millisecond)
			continue
		}
		if response.StatusCode != http.StatusSwitchingProtocols {
			t.Error("WebSocket connection not upgraded")
		}
		return connection
	}

}

func createServer(host string, port int) *WebSocketServer {
	webSocketServer := NewWebSocketServer(host, port)
	go webSocketServer.Start()

	return webSocketServer
}

// TestServer tests WebSocket communication between the server and a client
func TestServer(t *testing.T) {
	const host = "localhost"
	const port = 8000

	webSocketServer := createServer(host, port)
	defer webSocketServer.Shutdown()

	t.Run("test that server publishes events to connected clients", func(t *testing.T) {

		client := createClient(host, port, t)
		defer client.Close()

		publishedEvents := make(chan *feed.Event)
		go func() {
			receivedEvent := feed.Event{}
			err := client.ReadJSON(&receivedEvent)
			if err != nil {
				t.Errorf("Failed to read JSON from WebSocket connection: %s", err)
			}
			publishedEvents <- &receivedEvent
		}()

		event := createEvent()
		webSocketServer.Publish(event)

		select {
		case receivedEvent := <-publishedEvents:
			if receivedEvent.Type != event.Type && receivedEvent.Location != event.Location {
				t.Errorf("Received event %s does not match expected event %s", receivedEvent, event)
			}
		case <-time.After(5 * time.Second):
			t.Errorf("Did not receive event in 5 seconds")
		}

	})

}
