package datastructures

import "fmt"

//Vertex a struct example
type Vertex struct {
	X int
	Y int
}

//StructExample example for struct
func StructExample() {
	fmt.Println(Vertex{1, 2})
}

//reference: https://tour.golang.org/
