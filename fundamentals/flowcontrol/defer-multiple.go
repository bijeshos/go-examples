package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		//these defers will stack up and execute in last-in-first-out fashion
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

//reference: https://tour.golang.org/
