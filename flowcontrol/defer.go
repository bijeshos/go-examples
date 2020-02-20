package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

//reference: https://tour.golang.org/
