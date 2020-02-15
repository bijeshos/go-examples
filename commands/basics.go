package commands

import (
	"fmt",
	"time",
	"math/rand"
)

//PrintMessage prints a message
func PrintMessage(message string) {
	fmt.Println("Hello from main")
}

func PrintCurrentTime(){
	fmt.Println("Current time is: ", time.Now())
}

func PrintRandomNumber(){
	fmt.Println("A random number is: ", rand.Intn(10))
}