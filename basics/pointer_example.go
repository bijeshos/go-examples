package basics

import "fmt"

//Vertex a struct example
type Vertex struct {
	X int
	Y int
}

//PointerExample example to domonstrate pointers
func PointerExample() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

//StructExample example for struct
func StructExample() {
	fmt.Println(Vertex{1, 2})
}
