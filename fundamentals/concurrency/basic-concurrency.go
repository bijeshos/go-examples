package main

import "fmt"

func print(startValue int, increment int, count int) {
	fmt.Println("count: ", count, ", increment: ", increment)
	var currentValue = startValue

	for i := 0; i < count; i++ {

		fmt.Println("currentValue: : ", currentValue)
		currentValue = currentValue + increment
	}
}

func main() {
	go print(1, 1, 10)
	var input string
	fmt.Scanln(&input)
	//go print(50, 2, 10)
}
