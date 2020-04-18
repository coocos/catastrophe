package main

import (
	"flag"
	"os"
	"os/signal"
	"time"

	"github.com/coocos/catastrophe/internal/feed"
	"github.com/coocos/catastrophe/internal/server"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func configureLogger() {
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

func main() {

	configureLogger()

	host := flag.String("host", "localhost", "host to use")
	port := flag.Int("port", 8000, "port to use")
	flag.Parse()

	webSocketServer := server.NewWebSocketServer(*host, *port)

	stop := make(chan bool)
	events := feed.PollEvents(feed.NewClient(), time.NewTicker(60*time.Second), stop)

	// Stop the server gracefully on SIGINT
	go func() {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)

		for {
			select {
			case <-interrupt:
				stop <- true
				webSocketServer.Shutdown()
				log.Info("Shutting down")
				break
			case event := <-events:
				webSocketServer.Publish(event)
			}
		}

	}()

	webSocketServer.Start()

}
