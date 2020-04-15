package feed

import (
	"testing"
)

// TestParse tests that the parser can parse the event feed
func TestParse(t *testing.T) {

	feed := `<?xml version="1.0" encoding="ISO-8859-1"?>
    <rss version="2.0">
    <channel>
    <title>Pelastustoimen mediapalvelu: ensitiedotteet</title>
    <link>http://www.peto-media.fi</link>
    <description>Ensitiedotteita Suomesta....</description>
    <language>FI-fi</language>
    <item>
    <title>Liperi/Liperi, öljyvah./ymp.onnet. maalla:  pieni</title>
    <description>07.08.2019 17:17:04 Liperi/Liperi öljyvah./ymp.onnet. maalla:  pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 17:17:04 +0200</pubDate>
    </item>
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
    <item>
    <title>Salo/Salo, tieliikenneonnettomuus: pieni</title>
    <description>07.08.2019 17:06:26 Salo/Salo tieliikenneonnettomuus: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 17:06:26 +0200</pubDate>
    </item>
    </channel>
    </rss>`

	events, err := Parse(feed)
	if err != nil {
		t.Errorf("Failed to parse feed due to: %s", err)
	}
	if len(events) != 4 {
		t.Errorf("Feed should contain %d items, contained %d", 4, len(events))
	}
	event := events[0]
	if event.Location != "Salo" {
		t.Errorf("Event location is %s but should be %s", event.Location, "Pori")
	}
	expectedTime := "2019-08-07 17:06:26 +0200 +0200"
	if event.Time.String() != expectedTime {
		t.Errorf("Event time is %s but should be %s", event.Time.String(), expectedTime)
	}

}
