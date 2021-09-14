package main

import "fmt"

func main() {
	Fork1 := newFork()

	Fork1.in <- 500
	inUse := <-Fork1.out

	//msg := <-Fork1.in
	fmt.Println("Expected", Fork1.inUse, "Actual", inUse)
}
