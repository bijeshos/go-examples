package variables

import (
	"fmt"
)

//PrintDefaultVariableValues prints default values of variables
func PrintDefaultVariableValues() {
	var a bool
	var i int
	var f float32
	var c complex64

	const constant = "PI"

	fmt.Println("boolean:", a, ", int:", i, ", float: ", f, ", complex: ", c)
	fmt.Println("constant:", constant)
}
