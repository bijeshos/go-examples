package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	//splitlist splits a list of paths joined by OS specific separator (: or ;)
	fmt.Println(filepath.SplitList("/a/b/c:/usr/bin"))
}

//reference: https://golang.org/pkg
