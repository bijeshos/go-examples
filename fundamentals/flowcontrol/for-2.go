package main

import "fmt"

func main() {
	sum := 1
	// for without initialization and increment
	//for ; sum < 1000; {
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

//reference: https://tour.golang.org/