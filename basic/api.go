package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
)

func httpGet() {
	resp, err := http.Get("https://www.google.com/")
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func httpPost() {
	resp, err := http.Post("https://www.google.com/",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=test"))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func main() {
	httpGet()
	httpPost()
}

// https://ithelp.ithome.com.tw/articles/10159383