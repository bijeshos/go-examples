package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

}

//printMessage prints a message
func printMessage(message string) {
	fmt.Println("Printing message: ", message)
}

//printCurrentTime prints current time
func printCurrentTime() {
	fmt.Println("Current time is: ", time.Now())
}

//printRandomNumber prints a random number
func printRandomNumber() {
	fmt.Println("A random number is 2: ", rand.Intn(10))
}

//reference: https://tour.golang.org/
