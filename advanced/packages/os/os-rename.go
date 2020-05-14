package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	oldpath := "/home/bos/tmp/go-test/a.txt"
	newpath := "/home/bos/tmp/go-test/a-mod.txt"

	err := os.Rename(oldpath,newpath)
	if err != nil{
		log.Fatal(err)
	}else{
		fmt.Println("file renamed")
	}

	//os.Create("~/tmp/x.txt")
}

//reference: https://golang.org/pkg
