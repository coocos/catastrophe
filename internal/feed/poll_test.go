package feed

import (
	"reflect"
	"testing"
	"time"
)

type TestEventClient struct {
	events []Event
}

func (c *TestEventClient) LatestEvents() ([]Event, error) {
	return c.events[:len(c.events)-1], nil
}

func (c *TestEventClient) EventsSince(since time.Time) ([]Event, error) {

	eventsSince := []Event{}
	for _, event := range c.events {
		if event.Time.After(since) {
			eventsSince = append(eventsSince, event)
		}

	}
	return eventsSince, nil
}

func createEventClient(events []Event) *TestEventClient {
	return &TestEventClient{events}
}

// TestEventPolling tests polling for new events
func TestEventPolling(t *testing.T) {

	testEvents := []Event{
		{
			Location: "Helsinki",
			Type:     "tulipalo",
			Time:     time.Date(2020, time.June, 1, 12, 0, 0, 0, time.UTC),
		},
		{
			Location: "Oulu",
			Type:     "tieliikenneonnettomuus",
			Time:     time.Date(2020, time.June, 1, 12, 30, 0, 0, time.UTC),
		},
		{
			Location: "Espoo",
			Type:     "tulipalo",
			Time:     time.Date(2020, time.June, 1, 12, 45, 0, 0, time.UTC),
		},
	}

	t.Run("test that initial event and events after it are received", func(t *testing.T) {

		done := make(chan bool)
		events := PollEvents(createEventClient(testEvents), time.NewTicker(1*time.Microsecond), done)

		receivedEvents := []Event{}

		for event := range events {
			receivedEvents = append(receivedEvents, *event)
			if event.Location == "Espoo" {
				done <- true
			}
		}

		if !reflect.DeepEqual(receivedEvents, testEvents[1:]) {
			t.Errorf("Expected %v events but received %v", testEvents[1:], receivedEvents)
		}
	})

}
