package basics

import (
	"fmt"
	"math/rand"
	"time"
)

//PrintMessage prints a message
func PrintMessage(message string) {
	fmt.Println("Printing message: ", message)
}

//PrintCurrentTime prints current time
func PrintCurrentTime() {
	fmt.Println("Current time is: ", time.Now())
}

//PrintRandomNumber prints a random number
func PrintRandomNumber() {
	fmt.Println("A random number is 2: ", rand.Intn(10))
}

//reference: https://tour.golang.org/
