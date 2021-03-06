package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	//Ext returns the extension of the file; returns empty if no extension present

	fmt.Printf("No dots: %q\n", filepath.Ext("index"))
	fmt.Printf("One dot: %q\n", filepath.Ext("index.js"))
	fmt.Printf("Two dots: %q\n", filepath.Ext("main.test.js"))
}

//reference: https://golang.org/pkg
