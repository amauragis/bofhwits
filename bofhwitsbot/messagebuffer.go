package bofhwitsbot

import (
	"container/ring"
	"fmt"
	"github.com/thoj/go-ircevent"
)

var eventRing *ring.Ring

func MinIntS(v []int) (min int, idx int) {
	if len(v) > 0 {
		m = v[0]
	}
	for i := 1; i < len(v); i++ {
		if v[i] < m {
			idx = i
			m = v[i]
		}
	}
	return
}

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

func searchHistory(history []irc.Event, searchStr string) (match int) {

	const costs = LevenshteinCost{1, 2, 2} // set delete, insert, and substitution costs
	const timestampCostBuffer = 7

	var distances [len(history)]int

	for index := range history {
		distances[index] = Levenshtein(searchStr, "<"+history[index].Nick+"> "+history[index].Message(), &costs)

		fmt.Printf("%v: %v (%v) | %v\n", index, searchStr, "<"+history[index].Nick+"> "+history[index].Message(), dist)
	}
	min, minIdx = MinIntS(distances)

	// if the match is worse than half the length + the timestamp buffer, its probably trash
	if min > len(history[minIdx])/2+timestampCostBuffer {
		match = -1
	} else {
		match = minIdx
	}

	return match
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
