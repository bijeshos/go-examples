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

//Scale method pointer example
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//MethodPointerExample method pointer example
func MethodPointerExample() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

//ScaleFunc method indirection example
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//MethodIndirectionExample method indirection example
func MethodIndirectionExample() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}

//AbsFunc method indirection example
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//MethodIndirectionExample2 method indirection example
func MethodIndirectionExample2() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
