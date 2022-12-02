package main

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	errRequestFiled = errors.New("Request Failed With Unknown Issue")
)

func main() {
	var results = make(map[string]string)
	urls := []string{"https://google.com", "https://youtube.com", "https://facebook.com", "https://instagram.com", "https://github.com", "https://coupang.com", "https://airbnb.com"}
	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking : ", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if err == nil && resp.StatusCode >= 400 {
		fmt.Println(url, " -> Status Code : ", resp.StatusCode)
		return errRequestFiled
	}
	return nil
}
