package main

import "fmt"

func printFork() {
	fmt.Println("I am a fork")
}

type Fork struct {
	inUse     bool
	timesUsed int
	in        chan (bool)
	out       chan (bool)
}

func newFork() *Fork {
	in := make(chan bool)
	out := make(chan bool)

	f := Fork{inUse: false, timesUsed: 1, in: in, out: out}

	return &f
}

func occupy(f *Fork) {
	f.in <- true
}

func unoccupy(f *Fork) {
	f.in <- false
}

func activateChannels(f *Fork) {
	message := <-f.in
	fmt.Println(message, "hej")

	go activateChannels(f)
}
