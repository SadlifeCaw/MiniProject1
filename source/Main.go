package main

import (
	"fmt"
	"time"
)

func main() {

	//init forks
	var forks = make([]*fork, 5)
	for j := 0; j < 5; j++ {
		forks[j] = newFork(j)

		go activateForkChannels(forks[j])
	}

	//init philosophers
	var philosophers = make([]*philosopher, 5)
	for j := 0; j < 5; j++ {
		if j == 4 {
			philosophers[j] = newPhil(j, forks[(j+1)%5], forks[j]) //make 1 philosopher go opposite direction to avoid deadlocks
		} else {
			philosophers[j] = newPhil(j, forks[j], forks[(j+1)%5])
		}

		go eat(philosophers[j])
		go activatePhilosopherChannels(philosophers[j])
	}

	for {
		const delayInSeconds int32 = 6 //this time changes the forks in use / philosophers eating each query
		time.Sleep(time.Millisecond * time.Duration(delayInSeconds*1000))
		fmt.Println("Sending queries (every", delayInSeconds, "seconds):")

		//query forks
		for i := 0; i < 5; i++ {
			forks[i].in <- -1           //query for fork being in use
			inUse := <-forks[i].out     //fetch
			forks[i].in <- -2           //query for times used
			timesUsed := <-forks[i].out //fetch

			var inUseString string
			if inUse == 1 {
				inUseString = "Being used"
			} else {
				inUseString = "Not being used"
			}

			fmt.Println("Fork", i+1, ":", inUseString, "and has been used:", timesUsed, "times")
		}

		//query philosophers
		for i := 0; i < 5; i++ {

			philosophers[i].in <- -1        //query for eating/thinking
			eating := <-philosophers[i].out //fetch
			philosophers[i].in <- -2        //query for times eaten
			bites := <-philosophers[i].out  //fetch

			var eatingString string
			if eating == 1 {
				eatingString = "Eating"
			} else {
				eatingString = "Thinking"
			}

			fmt.Println("Philosopher", i+1, ":", eatingString, "and has eaten:", bites, "bites")
		}
	}
}
