package server

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/coocos/catastrophe/feed"
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
	connection, response, err := websocket.DefaultDialer.Dial(url.String(), nil)
	if err != nil {
		t.Errorf("Failed to connect to server %s", err)
	}
	if response.StatusCode != http.StatusSwitchingProtocols {
		t.Error("WebSocket connection not upgraded")
	}

	return connection
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

	publishedEvents := make(chan *feed.Event)

	client := createClient(host, port, t)
	defer client.Close()

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

}
