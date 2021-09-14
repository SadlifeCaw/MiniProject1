package main

import (
	"fmt"
	"time"
)

func main() {

	var forks = make([]*fork, 5)
	for j := 0; j < 5; j++ {
		var thisfork *fork = new(fork)
		thisfork.id = j
		forks[j] = new(fork)
	}
	var philosophers = make([]*philosopher, 5)
	for j := 0; j < 5; j++ {
		var phil *philosopher = new(philosopher)
		phil.id = j
		phil.leftFork = forks[j]
		phil.rightFork = forks[(j+1)%5]
		philosophers = append(philosophers, phil)
		go phil.eat()
	}
	fmt.Scanln()
	go reader()
}

func reader(){
	time.Sleep(time.Millisecond * 5000)
	var request string
	fmt.Scan(&request)
	//Implement different requests
}