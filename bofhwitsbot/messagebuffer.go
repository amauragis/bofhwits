package main

import (
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

type RingBuffer struct {
	inputChannel  <-chan Event
	outputChannel chan Event
}

func NewRingBuffer(inputChannel <-chan Event, outputChannel chan Event) *RingBuffer {
	return &RingBuffer{inputChannel, outputChannel}
}

func (r *RingBuffer) Run() {
	for v := range r.inputChannel {
		select {
		case r.outputChannel <- v:
		default:
			<-r.outputChannel
			r.outputChannel <- v
		}
	}
	close(r.outputChannel)
}

func main() {
	in := make(chan Event)
	out := make(chan Event, 5)
	rb := NewRingBuffer(in, out)
	go rb.Run()

	for i := 0; i < 10; i++ {
		in <- Event{"code", "raw" + fmt.Sprintf("%v", i), "nick", "host", "source", "user", []string{"Args1", "args2"}}
	}

	for res := range out {
		fmt.Println(res)
	}

	close(in)
}
