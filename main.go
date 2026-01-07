package main

import (
	"flag"
	"fmt"
)

var verbosePtr *bool
var userPtr *string
var aPtr *string

func main() {

	verbosePtr = flag.Bool("v", false, "verbose mode\n")
	userPtr = flag.String("u", "", "user\n")
	aPtr = flag.String("U", "", "user wordlist\n")
	flag.Parse()

	params := make(map[string]string)

	params["url"] = "http://localhost"

	makeRequest(params)
}

func makeRequest(params map[string]string) {
	// client := &http.Client{}

	if *aPtr != "" {
		fmt.Println(*aPtr)
	}

	if *userPtr != "" {
		fmt.Println(*userPtr)
	}

	fmt.Println(params["url"])

	// req, err := http.NewRequest("GET", params["url"], nil)
}
