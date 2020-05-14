package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//Though the package imported is math/rand, we can use rand directly.
	fmt.Println("My favorite number is", rand.Intn(10))
}

/*
A basic program to demonstrate package main and import packages
*/
//reference: https://tour.golang.org/
