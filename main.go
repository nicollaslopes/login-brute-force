package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var verbosePtr *bool
var userPtr *string
var UserPtr *string
var passwordPtr *string
var PasswordPtr *string

func main() {

	verbosePtr = flag.Bool("v", false, "verbose mode\n")
	userPtr = flag.String("u", "", "user\n")
	UserPtr = flag.String("U", "", "user wordlist\n")
	passwordPtr = flag.String("p", "", "password\n")
	PasswordPtr = flag.String("P", "", "password wordlist\n")
	flag.Parse()

	params := make(map[string]string)

	params["url"] = "http://localhost:9000/teste.php"

	makeRequest(params)
}

func makeRequest(params map[string]string) {

	data := url.Values{}
	data.Set("username", "admin")
	data.Set("password", "123")

	client := &http.Client{}
	req, err := http.NewRequest("POST", params["url"], strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(string(body))

	isFound := !strings.Contains(string(body), "incorrect")

	if isFound {
		fmt.Println("Found!")
	}

}
