package main

import "fmt"

func main() {
	//this will get executed after main completes execution
	defer fmt.Println("world")

	fmt.Println("hello")
}

//reference: https://tour.golang.org/
