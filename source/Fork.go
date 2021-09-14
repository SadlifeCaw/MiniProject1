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

	go readInput(&f)

	return &f
}

func readInput(f *fork) {
	for {
		message := <-f.in

		switch message {
		case 500: //set in use to true
			f.inUse = 500
			f.out <- f.inUse
		case 501: //set in use to false
			f.inUse = 501
			f.out <- f.inUse
		case 1:
			f.timesUsed++
			f.out <- f.timesUsed
		case 2:
			f.out <- f.timesUsed
		default:
			fmt.Println("Unknown")
		}
	}
}
