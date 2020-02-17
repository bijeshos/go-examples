package interfaces

//I interface implementaion
type I interface {
	M()
}

//T interface implementaion
type T struct {
	S string
}

//M interface implementaion
func (t T) M() {
	fmt.Println(t.S)
}

//InterfaceExample interface example
func InterfaceExample() {
	var i I = T{"hello"}
	i.M()
}
