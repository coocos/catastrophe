package main

import (
	"flag"
	"os"
	"os/signal"
	"time"

	"github.com/coocos/catastrophe/feed"
	"github.com/coocos/catastrophe/server"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func streamEventsToServer(eventStream <-chan *feed.Event, server server.EventServer) {
	for event := range eventStream {
		server.Publish(event)
	}
}

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
	defer webSocketServer.Start()

	// Stop the server gracefully on SIGINT
	go func() {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)
		<-interrupt
		webSocketServer.Shutdown()
	}()

	eventStream := make(chan *feed.Event)
	go feed.PollEvents(feed.NewRescueServiceClient(), time.NewTicker(60*time.Second), eventStream)
	go streamEventsToServer(eventStream, webSocketServer)

}
