package main

import (
	"fmt"
)

func printFork() {
	fmt.Println("I am a fork")
}

type Fork struct {
	inUse     int
	timesUsed int
	in        chan (int)
	out       chan (int)
}

func newFork() *Fork {
	in := make(chan int)
	out := make(chan int)

	f := Fork{inUse: 0, timesUsed: 0, in: in, out: out}

	go readInput(&f)

	return &f
}

func readInput(f *Fork) {
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
		default:
			fmt.Println("Unknown")
		}
	}
}
