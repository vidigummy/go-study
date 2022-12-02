package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://csbroker.io/problem"

func main() {
	// totalPages := getPages()
	// fmt.Println("totalPages : ", totalPages)
	for i := 0; i < 7; i++ {
		getPage(i)
	}
}

func getPage(page int) {
	pageURL := baseURL + "?page=" + strconv.Itoa(page+1)
	fmt.Println(pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find("a")
	searchCards.Each(func(i int, s *goquery.Selection) {

	})

}

// func getPages() int {
// 	pages := 0
// 	res, err := http.Get(baseURL)
// 	checkErr(err)
// 	checkCode(res)

// 	defer res.Body.Close()

// 	doc, err := goquery.NewDocumentFromReader(res.Body)
// 	checkErr(err)
// 	doc.Find(".ntcg5*").Each(func(i int, s *goquery.Selection) {
// 		pages = s.Find("button").Length()
// 	})
// 	return pages - 4
// }

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status: ", res.StatusCode)
	}
}
