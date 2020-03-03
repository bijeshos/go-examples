package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	//Match reports whether the name matches file pattern
	//returns boolean and error
	fmt.Println(filepath.Match("/home/catch/*", "/home/catch/foo"))
	fmt.Println(filepath.Match("/home/catch/*", "/home/catch/foo/bar"))
	fmt.Println(filepath.Match("/home/?opher", "/home/gopher"))
	fmt.Println(filepath.Match("/home/\\*", "/home/*"))

}
