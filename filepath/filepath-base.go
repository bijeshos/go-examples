package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	srcDir := "~/tmp"

	//Base returns the last element of the path

	fmt.Println(filepath.Base(srcDir))

	fmt.Println(filepath.Base("/foo/bar/baz.js"))
	fmt.Println(filepath.Base("/foo/bar/baz"))
	fmt.Println(filepath.Base("/foo/bar/baz/"))
	fmt.Println(filepath.Base("dev.txt"))
	fmt.Println(filepath.Base("../todo.txt"))
	fmt.Println(filepath.Base(".."))
	fmt.Println(filepath.Base("."))
	fmt.Println(filepath.Base("/"))
	fmt.Println(filepath.Base(""))

}

//reference: https://golang.org/pkg
