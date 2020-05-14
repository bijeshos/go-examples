package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	//Join joins any number of path elements to a single path

	elements := [][]string{
		{"a", "b", "c"},
		{"a", "b/c", ""},
		{"a/b", "c", ""},
		{"a/b", "/c", ""},
	}
	for _, row := range elements {
		joinedpath := filepath.Join(row[0], row[1], row[2])
		fmt.Printf("input: %q | joined-path: %q\n", row, joinedpath)
	}
}

//reference: https://golang.org/pkg
