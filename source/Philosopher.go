package main

import (
	//"fmt"
	"time"
)

type philosopher struct {
	id                  int
	bites               int
	eating              int
	leftFork, rightFork *fork
}

func (p philosopher) eat() {
	for true {
		p.leftFork.Lock()
		p.leftFork.beingUsed = 1
		p.rightFork.Lock()
		p.rightFork.beingUsed = 1
		p.eating = 1
		time.Sleep(time.Millisecond * 750)
		
		p.bites++
		p.eating = 0

		p.rightFork.Unlock()
		p.rightFork.used++
		p.rightFork.beingUsed = 0
		p.leftFork.Unlock()
		p.leftFork.used++
		p.leftFork.beingUsed = 0

		//fmt.Println("Philie ", p.id, " ", p.bites, " bites")
	}
}
