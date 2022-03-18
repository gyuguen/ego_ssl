package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("https://json-schema.org/draft/2020-12/schema")
	if err != nil {
		fmt.Println(fmt.Errorf("failed get https url: %w", err))
		panic(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("failed read body: %w", err))
		panic(err)
	}
	fmt.Println(string(body))
}
