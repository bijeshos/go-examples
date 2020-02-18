package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("Random number is: ", getRandomNumber())
	fmt.Println("Value of Pi is: ", getPi())

	fmt.Println("Sum of 10 and 20 is: ", add(10, 20))
	a, b := swap(50, 60)
	fmt.Println("Input: 50,60 | Output: ", a, b)
}

//getRandomNumber returns a random number
func getRandomNumber() int {
	return rand.Int()
}

//getPi returns value of PI
func getPi() float64 {
	return math.Pi
}

//add adds two values and returns the sum
func add(x int, y int) int {
	return x + y
}

//swap swaps two numbers
func swap(x, y int) (int, int) {
	return y, x
}

//reference: https://tour.golang.org/
