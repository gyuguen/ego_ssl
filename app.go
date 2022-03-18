package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	client := http.Client{
		Transport: &http.Transport{TLSClientConfig: tlsConfig}}
	res, err := client.Get("https://json-schema.org/draft/2020-12/schema")
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
