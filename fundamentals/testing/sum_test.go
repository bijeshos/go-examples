package main

import (
	"fmt"
	"testing"
)

//unit test
func TestSum(t *testing.T) {
	total := Sum(2, 4)
	if total != 6 {
		t.Errorf("sum was incorrect, expected: %d, actual: %d", 6, total)
	}
}

// how to execute the test:
// open terminal, cd to this directory and execute the following
// go test
// go test -cover # to know the coverage as well

// to visualize coverage
//go test -cover -coverprofile=c.out
//go tool cover -html=c.out -o coverage.html


func BenchmarkSum(b *testing.B) {
	total := Sum(2, 4)
	fmt.Println("total:", total)
}

//to perform benchmarking
//go test -bench=.
