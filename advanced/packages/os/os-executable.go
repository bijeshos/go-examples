package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	execname, err := os.Executable()
	if err != nil{
		log.Fatal(err)
	}else{
		fmt.Println("executable-name:",execname)
	}

}

//reference: https://golang.org/pkg
