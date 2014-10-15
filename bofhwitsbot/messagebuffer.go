package main

import (
	"container/ring"
	"fmt"
)

// A struct to represent an event.
type Event struct {
	Code      string
	Raw       string
	Nick      string //<nick>
	Host      string //<nick>!<usr>@<host>
	Source    string //<host>
	User      string //<usr>
	Arguments []string
}

var eventRing *ring.Ring

func InitRing(size int) {
	eventRing = ring.New(size)
}

func addEvent(newEvent Event) {
	// move backward in the ring so we're at the oldest entry
	eventRing = eventRing.Prev()
	eventRing.Value = newEvent
}

func getHistory() []Event {
	var history []Event
	history = []Event{}
	eventRing.Do(func(x interface{}) {
		if x != nil {
			history = append(history, x.(Event))
		}
	})

	return history
}

func main() {
	InitRing(15)
	for i := 1; i <= 25; i++ {
		addEvent(Event{"code", "raw" + fmt.Sprintf("%v", i), "nick", "host", "source", "user", []string{"Args1", "args2"}})
	}

	history := getHistory()
	for i := range history {
		fmt.Printf("%v: %v\n", i, history[i])
	}
}
