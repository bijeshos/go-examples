package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	//Join joins any number of path elements to a single path
	fmt.Println(filepath.Join("a", "b", "c"))
	fmt.Println(filepath.Join("a", "b/c"))
	fmt.Println(filepath.Join("a/b", "c"))
	fmt.Println(filepath.Join("a/b", "/c"))
}
