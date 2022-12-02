package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type requestResult struct {
	url    string
	status string
}

func main() {
	results := make(map[string]string)
	urls := []string{"https://google.com", "https://youtube.com", "https://facebook.com", "https://instagram.com", "https://github.com", "https://coupang.com", "https://airbnb.com"}
	resultChannel := make(chan requestResult)
	for _, url := range urls {
		go hitURL(url, resultChannel)
	}
	for i := 0; i < len(urls); i++ {
		// fmt.Println(<-resultChannel)
		result := <-resultChannel
		results[result.url] = result.status
	}
	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan requestResult) {
	// fmt.Println("Checking ", url)
	resp, err := http.Get(url)
	result_status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		result_status = "FAILED " + strconv.Itoa(resp.StatusCode)
	}
	c <- requestResult{url: url, status: result_status}

}
