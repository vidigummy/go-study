package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type result struct {
	url    string
	status string
}

func main() {
	urls := []string{"https://google.com", "https://youtube.com", "https://facebook.com", "https://instagram.com", "https://github.com", "https://coupang.com", "https://airbnb.com"}
	resultChannel := make(chan result)
	for _, url := range urls {
		go hitURL(url, resultChannel)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-resultChannel)
	}
}

func hitURL(url string, c chan<- result) {
	// fmt.Println("Checking ", url)
	resp, err := http.Get(url)
	result_status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		result_status = "FAILED " + strconv.Itoa(resp.StatusCode)
	}
	c <- result{url: url, status: result_status}

}
