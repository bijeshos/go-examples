package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	//Dir returns the last directory in the path, without the last slashes

	paths := []string{
		"/foo/bar/baz.js",
		"/foo/bar/baz",
		"/foo/bar/baz/",
		"/dirty//path///",
		"dev.txt",
		"../todo.txt",
		"..",
		".",
		"/",
		"",
	}
	for _, path := range paths {
		dir := filepath.Dir(path)
		fmt.Printf("input: %q | dir: %q\n", path, dir)
	}

}

//reference: https://golang.org/pkg
