package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Request struct {
	Method  string
	URI     string
	Headers map[string]string
	Body    string
}

func helpInfo() {
	fmt.Println("Request URI")
	fmt.Println("gourl [COMMAND]")
	fmt.Println("COMMAND list:")
	fmt.Println("\t--method [post|get], -m:\t\trequest method, eg: POST, GET")
	fmt.Println("\t--body [string], -b, -d:\t\tpost request body")
	fmt.Println("\t--header [string], -h, -H\t\trequest header")
	fmt.Println("\t--json, -j\t\tequals to -H 'Content-Type: application/json'")
	fmt.Println("\t--uri, optional\t\trequest uri\n")
	fmt.Println("Examples:")
	fmt.Println("\tgourl http://example.com -d '{\"content\": \"hhh\"}' -j")
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		helpInfo()
		return
	}
	request := &Request{Headers: make(map[string]string)}
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-m":
			fallthrough
		case "--method":
			request.Method = args[i+1]
			i++
		case "-d":
			fallthrough
		case "-b":
			fallthrough
		case "--data-binary":
			fallthrough
		case "--body":
			request.Body = args[i+1]
			i++
		case "-h":
			fallthrough
		case "-H":
			fallthrough
		case "--header":
			headerKV := strings.Split(args[i+1], ":")
			request.Headers[headerKV[0]] = strings.Join(headerKV[1:], "")
			i++
		case "--uri":
			i++
			fallthrough
		case "-j":
			fallthrough
		case "--json":
			request.Headers["Content-Type"] = "application/json"
		default:
			if !strings.HasPrefix(args[i], "-") {
				if !strings.HasPrefix(args[i], "http") {
					request.URI = fmt.Sprintf("http://%s", args[i])
				} else {
					request.URI = args[i]
				}
			}
		}
	}
	var requestBody io.Reader
	if request.Body != "" {
		request.Method = "POST"
		requestBody = strings.NewReader(request.Body)
	}
	req, err := http.NewRequest(strings.ToUpper(request.Method), request.URI, requestBody)
	// fmt.Println(request)
	// fmt.Println(strings.Join(args, ","))
	for k, v := range request.Headers {
		req.Header.Add(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// handle error
		fmt.Printf("error: %s", err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}
	var out bytes.Buffer
	err2 := json.Indent(&out, body, "", "  ")
	if err2 != nil {
		fmt.Printf("%s", body)
		return
	}
	fmt.Printf("%s", out.String())
}
