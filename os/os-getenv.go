package main

import (
	"fmt"
	"os"
)

func main() {
	//os.Setenv("NAME", "gopher")
	//os.Setenv("BURROW", "/usr/gopher")

	fmt.Printf("GOROOT: %s , \nGOPATH: %s \n", os.Getenv("GOROOT"), os.Getenv("GOPATH"))

}
