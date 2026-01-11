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
var passwordPtr *string
var fieldsLoginPtr *string
var valuesPtr *string
var extraFieldsPtr *string

func main() {

	verbosePtr = flag.Bool("v", false, "verbose mode\n")
	userPtr = flag.String("u", "", "user\n")
	passwordPtr = flag.String("p", "", "password\n")
	fieldsLoginPtr = flag.String("fields-login", "", "params\n")
	valuesPtr = flag.String("values", "", "values\n")
	extraFieldsPtr = flag.String("extra-fields", "", "values\n")
	flag.Parse()
	validateFlags()

	params := make(map[string]string)

	params["url"] = "http://localhost:9000/teste.php"

	//  -fields-login "username,password" -extra-fields="cookie"  -values="abc123" -msg "incorrect"
	makeRequest(params)
}

func makeRequest(params map[string]string) {

	// if *PasswordPtr != "" {
	// 	file, err := os.Open(*PasswordPtr)
	// 	if err != nil {
	// 		log.Fatalf("Error opening file: %s", err)
	// 	}

	// 	defer file.Close()

	// 	scanner := bufio.NewScanner(file)
	// 	for scanner.Scan() {
	// 		// line := scanner.Text()
	// 	}
	// }

	data := setParams()

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

	paramsTotal := len(strings.Split(*extraFieldsPtr, ","))
	valuesTotal := len(strings.Split(*valuesPtr, ","))

	if paramsTotal != valuesTotal {
		fmt.Println("Params does not match with values!")
		os.Exit(1)
	}

	if *userPtr == "" {
		fmt.Println("Error: -u or -U flag is required.")
		flag.Usage()
		os.Exit(1)
	}

	if *passwordPtr == "" {
		fmt.Println("Error: -p flag is required.")
		flag.Usage()
		os.Exit(1)
	}
}

func setParams() url.Values {

	fieldsLogin := strings.Split(*fieldsLoginPtr, ",")

	data := url.Values{}
	data.Set(fieldsLogin[0], *userPtr)
	data.Set(fieldsLogin[1], *passwordPtr)

	return data
}
