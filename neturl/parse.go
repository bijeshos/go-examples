package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	//parses rawurl into a URL structure.
	u, err := url.Parse("http://bing.com/search?q=dotnet")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)
	u.Scheme = "https"
	u.Host = "google.com"
	q := u.Query()
	q.Set("q", "golang")
	u.RawQuery = q.Encode()
	fmt.Println(u)
}