package main

import (
	"fmt"
	"time"
)

func main() {

	var forks = make([]*fork, 5)
	for j := 0; j < 5; j++ {
		forks[j] = newFork(j)
	}
	var philosophers = make([]*philosopher, 5)
	for j := 0; j < 5; j++ {
		philosophers[j] = newPhil(j, forks[j], forks[(j+1)%5])
	}
	
	for {
		fmt.Println("Sending queries:")
		for i := 0; i < 5; i++ {
			time.Sleep(time.Millisecond * 1464)
			forks[i].in <- -1
			inUse := <-forks[i].out
			var inUseString string
			if (inUse == 1){
				inUseString = "Being used"
			} else {
				inUseString = "Not being used"
			}
			forks[i].in <- -2
			timesUsed := <-forks[i].out
			fmt.Println("Fork", i+1, ":", inUseString, "and has been used:", timesUsed, "times.")
		} 
		time.Sleep(time.Millisecond * 2000)
		for i := 0; i < 5; i++ {
			time.Sleep(time.Millisecond * 1464)
			philosophers[i].in <- -1
			eating := <-philosophers[i].out
			var eatingString string
			if (eating == 1){
				eatingString = "Eating"
			} else {
				eatingString = "Thinking"
			}
			philosophers[i].in <- -2
			bites := <-philosophers[i].out
			fmt.Println("Philosopher", i+1, ":", eatingString, "and has eaten: ", bites, "bites.")			
		} 
	}
}
