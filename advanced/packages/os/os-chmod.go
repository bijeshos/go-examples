package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	err := os.Chmod("/home/bos/tmp/a.txt", 0644)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("changed mode")
	}

	execname, err := os.Executable()
	fmt.Println("execname:", execname, ", error: ", err)

	//func Chown(name string, uid, gid int) error //Chown changes the numeric uid and gid of the named file.
}

//reference: https://golang.org/pkg
