package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	//Match reports whether the name matches file pattern
	//returns boolean and error
	inputs := [][]string{
		{"/home/catch/*", "/home/catch/foo"},
		{"/home/catch/*", "/home/catch/foo/bar"},
		{"/home/?opher", "/home/gopher"},
		{"/home/\\*", "/home/*"},
	}
	for _, row := range inputs {
		match, _ := filepath.Match(row[0], row[1])
		fmt.Printf("pattern: %q | name: %q | match: %t\n", row[0], row[1], match)
	}

}

//reference: https://golang.org/pkg
