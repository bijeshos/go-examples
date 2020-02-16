package main

import (
	"fmt"

	"github.com/bijeshos/go.examples/basics"
	"github.com/bijeshos/go.examples/flowcontrol"
	"github.com/bijeshos/go.examples/variables"
)

func main() {
	basics.PrintMessage("test message")
	fmt.Println("Random number is: ", basics.GetRandomNumber())
	fmt.Println("Value of Pi is: ", basics.GetPi())

	fmt.Println("Sum of 10 and 20 is: ", basics.Add(10, 20))
	a, b := basics.Swap(50, 60)
	fmt.Println("Input: 50,60 | Output: ", a, b)

	variables.PrintDefaultVariableValues()
	flowcontrol.ForLoop()
	flowcontrol.Sqrt(8.0)
	flowcontrol.Pow(2, 2, 1)
	flowcontrol.Switch()
	flowcontrol.SwithWithNoCondition()

	basics.PointerExample()
	basics.StructExample()
}
