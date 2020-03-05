package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	//create a file, if it does not exit
	//truncates the file, if it exists
	//ensure that the go process has permissions to perform read/write actions
	file, err := os.Create("/home/bos/tmp/go-test/x.txt")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("file created; ", file.Name())
	}

}

//reference: https://golang.org/pkg
