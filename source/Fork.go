package main

import (
	"fmt"
	"sync"
)

type fork struct {
	sync.Mutex
	inUse     int
	timesUsed int
	in        chan (int)
	out       chan (int)
	id        int
}

func newFork(init int) *fork {
	in := make(chan int)
	out := make(chan int)

	f := fork{id: init,
		inUse:     0,
		timesUsed: 0,
		in:        in,
		out:       out,
	}

	return &f
}

func activateForkChannels(f *fork) {
	for {
		message := <-f.in

		switch message {
		case -1:
			f.out <- f.inUse
		case -2:
			f.out <- f.timesUsed
		default:
			fmt.Println("Unknown")
		}
	}
}
