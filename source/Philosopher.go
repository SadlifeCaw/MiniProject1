package main

import (
	"fmt"
	"math/rand"
	"time"
)

type philosopher struct {
	id                  int
	bites               int
	eating              int
	leftFork, rightFork *fork
	in                  chan (int)
	out                 chan (int)
}

func newPhil(init int, leftFork *fork, rightFork *fork) *philosopher {
	in := make(chan int)
	out := make(chan int)

	p := philosopher{
		id:        init,
		eating:    0,
		bites:     0,
		in:        in,
		out:       out,
		leftFork:  leftFork,
		rightFork: rightFork,
	}

	return &p
}

func eat(p *philosopher) {
	for {
		if ShouldIEat(p) {
			p.leftFork.Lock()
			p.rightFork.Lock()
			p.leftFork.inUse = 1
			p.rightFork.inUse = 1

			p.eating = 1

			time.Sleep(time.Millisecond * 400)

			p.bites++

			time.Sleep(time.Millisecond * 400)

			p.eating = 0

			p.rightFork.timesUsed++
			p.rightFork.inUse = 0

			p.leftFork.timesUsed++
			p.leftFork.inUse = 0

			p.rightFork.Unlock()
			p.leftFork.Unlock()
		}
	}
}

func ShouldIEat(p *philosopher) bool {
	randomNumber := rand.Intn(2)
	return randomNumber == 1
}

func activatePhilosopherChannels(p *philosopher) {
	for {
		message := <-p.in

		switch message {
		case -1:
			p.out <- p.eating
		case -2:
			p.out <- p.bites

		default:
			fmt.Println("Unknown")
		}
	}
}
