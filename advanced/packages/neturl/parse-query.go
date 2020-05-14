package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"
)

func main() {
	//parses the URL-encoded query string and returns a map listing the values specified for each key.
	m, err := url.ParseQuery(`x=1&y=2&y=3;z`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(toJSON(m))
}

func toJSON(m interface{}) string {
	js, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	return strings.ReplaceAll(string(js), ",", ", ")
}