package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	path, err := os.Getwd()
	fmt.Println("current working directory: ", path)
	//changes current working directory to named directory
	err = os.Chdir("/tmp")
	if err != nil {
		log.Fatal(err)
	}
	path, err = os.Getwd()
	fmt.Println("current working directory: ", path)

}
//reference: https://golang.org/pkg