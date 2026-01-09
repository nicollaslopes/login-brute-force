package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

var verbosePtr *bool
var userPtr *string
var UserPtr *string
var passwordPtr *string
var PasswordPtr *string
var paramsPtr *string
var valuesPtr *string

func main() {

	verbosePtr = flag.Bool("v", false, "verbose mode\n")
	userPtr = flag.String("u", "", "user\n")
	UserPtr = flag.String("U", "", "user wordlist\n")
	passwordPtr = flag.String("p", "", "password\n")
	PasswordPtr = flag.String("P", "", "password wordlist\n")
	paramsPtr = flag.String("params", "", "params\n")
	valuesPtr = flag.String("values", "", "values\n")
	flag.Parse()
	validateFlags()

	params := make(map[string]string)

	params["url"] = "http://localhost:9000/teste.php"

	//  -params "username,password,page" -values "^USER^,^PASS^,test" -msg "incorrect"
	// user=^USER^&password=^PASS^&submit=Login:Login Failed
	makeRequest(params)
}

func makeRequest(params map[string]string) {

	var user string

	paramsTotal := len(strings.Split(*paramsPtr, ","))
	valuesTotal := len(strings.Split(*valuesPtr, ","))

	if paramsTotal != valuesTotal {
		fmt.Println("Params does not match with values!")
		os.Exit(1)
	}

	if *userPtr != "" {
		user = *userPtr
	}

	if *UserPtr != "" {
		user = "test"
	}

	userParam := strings.Split(*paramsPtr, ",")

	data := url.Values{}
	data.Set(userParam[0], user)
	data.Set(userParam[1], "123")

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

func validateFlags() {
	if *userPtr != "" && *UserPtr != "" {
		fmt.Println("You can not use -u and -U flags together")
		os.Exit(1)
	}

	if *passwordPtr != "" && *PasswordPtr != "" {
		fmt.Println("You can not use -p and -P flags together")
		os.Exit(1)
	}
}
