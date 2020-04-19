package main

import (
	"flag"
	"fmt"
	"net/url"
	"sync"

	"github.com/coocos/catastrophe/internal/feed"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func configureLogger() {
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

func createClient(host string, port int) (*websocket.Conn, error) {
	url := url.URL{
		Scheme: "ws",
		Host:   fmt.Sprintf("%s:%d", host, port),
		Path:   "/websocket",
	}
	connection, _, err := websocket.DefaultDialer.Dial(url.String(), nil)
	if err != nil {
		log.Error("Failed to connect client - is the server running?")
		return nil, err
	}
	return connection, nil
}

func main() {

	configureLogger()

	host := flag.String("host", "localhost", "host to connect to")
	port := flag.Int("port", 8000, "port to connect to")
	count := flag.Int("count", 1, "how many clients to spawn")
	flag.Parse()

	wg := sync.WaitGroup{}
	wg.Add(*count)

	for i := 0; i < *count; i++ {

		log.WithFields(log.Fields{
			"count": i,
		}).Info("Spawning client")

		go func() {

			client, err := createClient(*host, *port)
			if err != nil {
				log.Fatal("Failed to connect client")
			}
			for {
				event := feed.Event{}
				err := client.ReadJSON(&event)
				if err != nil {
					log.WithFields(log.Fields{
						"error": err,
					}).Errorf("Failed to receive event")
					break
				} else {
					log.WithFields(log.Fields{
						"event": event,
					}).Info("Client received event")
				}
			}
			client.Close()
			wg.Done()
		}()
	}

	wg.Wait()

}
