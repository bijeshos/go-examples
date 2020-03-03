package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	paths := []string{
		"/home/bos/profile.jpg",
		"/mnt/photos/",
		"rabbit.jpg",
		"/usr/local//go",
	}
	for _, p := range paths {
		//split splits path immediately following last separator, separating path to directory and file name
		//if no separator in the path, returns empty dir and file name
		//
		dir, file := filepath.Split(p)
		fmt.Printf("input: %q\n\tdir: %q\n\tfile: %q\n", p, dir, file)
	}
}
