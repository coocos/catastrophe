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
	defer webSocketServer.Start()

	eventStream := make(chan *feed.Event)

	// Stop the server gracefully on SIGINT
	go func() {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)
		<-interrupt
		close(eventStream)
		webSocketServer.Shutdown()
	}()

	// TODO: Maybe this should return a channel instead? Then once you close the channel it's done?
	// So instead of passing eventStream, it internally creates a goroutine which puts events into the
	// channel it returns
	go feed.PollEvents(feed.NewRescueServiceClient(), time.NewTicker(60*time.Second), eventStream)

	go func() {
		for event := range eventStream {
			webSocketServer.Publish(event)
		}
	}()

}
