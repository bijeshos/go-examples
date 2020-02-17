package basics

import (
	"math"
	"math/rand"
)

//GetRandomNumber returns a random number
func GetRandomNumber() int {
	return rand.Int()
}

//GetPi returns value of PI
func GetPi() float64 {
	return math.Pi
}

//Add adds two values and returns the sum
func Add(x int, y int) int {
	return x + y
}

//Swap swaps two numbers
func Swap(x, y int) (int, int) {
	return y, x
}

//reference: https://tour.golang.org/
