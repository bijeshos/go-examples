package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u, err := url.Parse("http://example.com/path with spaces")
	if err != nil {
		log.Fatal(err)
	}
	//EscapedPath returns the escaped form of u.Path
	fmt.Println(u.EscapedPath())
}