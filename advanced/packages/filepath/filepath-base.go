package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	//Base returns the last element of the path

	paths := []string{
		"~/tmp",
		"/foo/bar/baz.js",
		"/foo/bar/baz",
		"/foo/bar/baz/",
		"dev.txt",
		"../todo.txt",
		"..",
		".",
		"/",
		"",
	}
	for _, path := range paths {
		base := filepath.Base(path)
		fmt.Printf("input: %q | base: %q\n", path, base)
	}

}

//reference: https://golang.org/pkg
