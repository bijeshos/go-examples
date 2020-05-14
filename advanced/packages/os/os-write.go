package main

import (
	"fmt"
	"log"
	"os"
)

func main(){

	file, err := os.Create("/home/bos/tmp/go-test/test-file-2.txt")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("file created; ", file.Name())
	}

	d2 := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}  //means, hello world
	n2, err := file.Write(d2)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	fmt.Println(n2, "bytes written successfully")
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}


}

