package main

import (
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
