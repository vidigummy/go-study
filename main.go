package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://search.naver.com/search.naver?where=nexearch&sm=tab_etc&qvt=0&query=%EC%98%81%ED%99%94%EC%98%88%EB%A7%A4%EC%82%AC%EC%9D%B4%ED%8A%B8"

func main() {
	totalPages := getPages()
	fmt.Println("totalPages : ", totalPages)
	for i := 0; i < totalPages; i++ {
		getPage(i)
	}
}

func getPage(page int) {
	pageURL := "https://search.naver.com/search.naver?display=15&f=&filetype=0&page=" + strconv.Itoa(page) + "&query=%EC%98%81%ED%99%94%EC%98%88%EB%A7%A4%EC%82%AC%EC%9D%B4%ED%8A%B8&qvt=0&research_url=&sm=tab_pge&start=31&where=web"
	fmt.Println(pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	cardDoc, err := goquery.NewDocumentFromReader(res.Body)
	searchCards := cardDoc.Find("a.link_tit")
	searchCards.Each(func(i int, s *goquery.Selection) {
		id, _ := s.Attr("href")
		fmt.Println(id)
	})
	fmt.Println(res.Body)

}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	doc.Find(".api_sc_page_wrap").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})
	return pages - 2
}

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
