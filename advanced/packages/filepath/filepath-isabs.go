package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	//IsAbs checks if the path is absolute
	paths := []string{
		"/home/gopher",
		".bashrc",
		"..",
		".",
		"/",
		"",
	}
	for _, path := range paths {
		isabs := filepath.IsAbs(path)
		fmt.Printf("input: %q | isabs: %t\n", path, isabs)
	}

}
