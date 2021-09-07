package main

//import "fmt"

func main() {
	Fork1 := newFork()

	go activateChannels(Fork1)

	occupy(Fork1)

	//msg := <-Fork1.in
	//fmt.Println(msg)
}
