package methods

import (
	"fmt"
	"math"
)

//Vertex used for holding vertex
type Vertex struct {
	X, Y float64
}

//Abs method example
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//MethodExample method example
func MethodExample() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

//Abs fuction example
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//FunctionExample function example
func FunctionExample() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
