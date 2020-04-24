# Catastrophe

🚒 WebSocket API for Finnish rescue service events

## API

The server will publish all new events as they happen to all WebSocket clients currently connected to `/websocket/`. Events are serialized in the following format using JSON:

```json
{
    "type": "palohälytys",
    "location": "Harjavalta",
    "time": "2020-04-15T18:57:00+02:00",
    "description": "15.04.2020 18:57:00 Harjavalta/Harjavalta palohälytys"
}
```

The latest event is always published to all WebSocket clients when they connect.

## How to use

Start the server:

```bash
go run cmd/server/server.go --host localhost --port 8000
```

Or use Docker:

```bash
docker build -t catastrophe .
docker run -d catastrophe
```

And then you can connect a WebSocket client using JavaScript for example:

```javascript
const socket = new WebSocket("ws://localhost:8000/websocket")
socket.onmessage = (msg) => console.log(msg);
```

## How to run tests

To run all tests:

```bash
go test ./...
```

## Motivation

The Finnish rescue services provide an [RSS feed](http://www.peto-media.fi/tiedotteet/rss.xml) but there is no simple event-driven API which yields JSON. Well, now there is! Catastrophe periodically polls the feed and publishes all new events as JSON to all connected WebSocket clients.
