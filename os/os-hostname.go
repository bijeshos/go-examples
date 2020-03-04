package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("hostname: %s \n", hostname)

}
