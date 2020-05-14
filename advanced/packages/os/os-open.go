package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//open file for read access
	file, err := os.Open(".gitignore")
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 1000)
	//read data into byte array and assign no. of bytes to count
	count, err := file.Read(data)
	defer file.Close() //closing the file after all operations are done

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("read %d bytes: %q\n", count, data[:count])

}

//reference: https://golang.org/pkg