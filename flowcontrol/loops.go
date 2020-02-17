package flowcontrol

import "fmt"

//ForLoop contains basic loop example
func ForLoop() {

	fmt.Println("executing for loop")
	//approach 1
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//approach 2
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	//approach 3
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

//reference: https://tour.golang.org/
