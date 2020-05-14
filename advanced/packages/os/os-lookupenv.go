package main

import (
	"fmt"
	"os"
)

func main() {

	printEnv := func(key string) {
		val, ok := os.LookupEnv(key)
		if !ok {
			fmt.Printf("%s not set\n", key)
		} else {
			fmt.Printf("%s=%s\n", key, val)
		}
	}

	printEnv("GOROOT")
	printEnv("GOPATH")
	printEnv("GOBIN")

}

//reference: https://golang.org/pkg
