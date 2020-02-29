package main

import "fmt"

func main() {
	sum := 0
	//simple for loop, no paranthesis needed
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

//reference: https://tour.golang.org/
