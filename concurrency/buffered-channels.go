package main

import "fmt"

func main() {
	// this channel has a buffer size of 2
	// sends are blocked when channel is full
	// receives are blocked when channel is empty
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

//reference: https://tour.golang.org/
