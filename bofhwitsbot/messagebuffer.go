package bofhwitsbot

import (
	"container/ring"
	"fmt"
	"github.com/thoj/go-ircevent"
)

var eventRing *ring.Ring

func InitRing(size int) {
	eventRing = ring.New(size)
}

func addEvent(newEvent *irc.Event) {
	// move backward in the ring so we're at the oldest entry
	eventRing = eventRing.Prev()
	eventRing.Value = *newEvent
}

func getHistory() []irc.Event {
	var history []irc.Event
	history = []irc.Event{}
	eventRing.Do(func(x interface{}) {
		if x != nil {
			history = append(history, x.(irc.Event))
		}
	})

	return history
}

func searchHistory(history []irc.Event, searchStr string) (matchIndicies []int) {

	costs := LevenshteinCost{1, 2, 2} // set delete, insert, and substitution costs
	for index := range history {
		dist := Levenshtein(searchStr, "<"+history[index].Nick+"> "+history[index].Message(), &costs)
		if dist < 6 {
			matchIndicies = append(matchIndicies, index)
		}
		fmt.Printf("%v: %v (%v) | %v\n", index, searchStr, "<"+history[index].Nick+"> "+history[index].Message(), dist)
	}

	return matchIndicies
}

// func main() {
// 	InitRing(15)
// 	for i := 1; i <= 25; i++ {
// 		addEvent(Event{"code", "raw" + fmt.Sprintf("%v", i), "nick", "host", "source", "user", []string{"Args1", "args2"}})
// 	}

// 	history := getHistory()
// 	for i := range history {
// 		fmt.Printf("%v: %v\n", i, history[i])
// 	}
// }
