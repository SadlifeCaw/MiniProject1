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

func newFork(id int) *fork {
	in := make(chan int)
	out := make(chan int)
	id = id

	f := fork{inUse: 0, timesUsed: 0, in: in, out: out}

	go readFork(&f)

	return &f
}

func readFork(f *fork) {
	for {
		message := <-f.in

		switch message {
		case -1: //set in use to true
			f.inUse = message
			f.out <- f.inUse
		case -2: //set in use to false
			f.inUse = message
			f.out <- f.inUse
		case -3:
			f.out <- f.timesUsed
		default:
			fmt.Println("Unknown")
		}
	}
}
