package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	path := "/home/bos/tmp/go-test/os-test/"

	err := os.RemoveAll(path)
	if err != nil{
		log.Fatal(err)
	}else{
		fmt.Println("directory removed")
	}
}

//reference: https://golang.org/pkg
