package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Create("/home/bos/tmp/go-test/test-file.txt")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("file created; ", file.Name())
	}

	l, err := file.WriteString("Hello World")
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

}
