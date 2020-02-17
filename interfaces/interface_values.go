package interfaces

import (
	"fmt"
	"math"
)

//I interface values example
type I interface {
	M()
}

//T interface values example
type T struct {
	S string
}

//M interface values example
func (t *T) M() {
	fmt.Println(t.S)
}

//F interface values example
type F float64

//M interface values example
func (f F) M() {
	fmt.Println(f)
}

//InterfaceValuesExample interface values example
func InterfaceValuesExample() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

//reference: https://tour.golang.org/
