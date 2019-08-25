package feed

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func NewTestClient(baseUrl string) *Client {
	return &Client{
		baseUrl + "/tiedotteet/rss.xml",
	}
}

func NewTestServer() *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := `<?xml version="1.0" encoding="ISO-8859-1"?>
    <rss version="2.0">
    <channel>
    <title>Pelastustoimen mediapalvelu: ensitiedotteet</title>
    <link>http://www.peto-media.fi</link>
    <description>Ensitiedotteita Suomesta....</description>
    <language>FI-fi</language>
    <item>
    <title>Mustasaari/Korsholm, maastopalo: pieni</title>
    <description>07.08.2019 17:13:04 Mustasaari/Korsholm maastopalo: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 17:13:04 +0200</pubDate>
    </item>
    <item>
    <title>Espoo/Esbo, tieliikenneonnettomuus: pieni</title>
    <description>07.08.2019 17:12:16 Espoo/Esbo tieliikenneonnettomuus: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 17:12:16 +0200</pubDate>
    </item>
    </channel>
    </rss>`
		fmt.Fprintln(w, response)
	}))
	return server
}

func TestClientLatestEvents(t *testing.T) {

	server := NewTestServer()
	defer server.Close()

	client := NewTestClient(server.URL)
	events, err := client.LatestEvents()
	if err != nil {
		t.Errorf("Failed to fetch latest events")
	}
	if len(events) != 2 {
		t.Errorf("Client returned %d events instead of %d", len(events), 3)
	}

}

func TestClientEventsSince(t *testing.T) {

	server := NewTestServer()
	defer server.Close()

	client := NewTestClient(server.URL)
	events, _ := client.LatestEvents()
	firstEvent := events[0]

	newEvents, _ := client.EventsSince(firstEvent.Time)

	if len(newEvents) != 1 {
		t.Errorf("Client returned %d new events instead of %d", len(newEvents), 1)
	}
	if newEvents[0].Time.Before(firstEvent.Time) {
		t.Errorf("Client returned a previous event")
	}

}
