package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	dirname := "."

	//the files returned are not in sorted order
	f, err := os.Open(dirname)
	if err != nil {
		log.Fatal(err)
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

//reference: https://golang.org/pkg
