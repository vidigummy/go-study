package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/github"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		userName := "dengoyoon"
		repoList, err := getRepo(userName)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(repoList)
		languageMap := make(map[string]int)
		repoChan := make(chan map[string]int)
		for _, repoName := range repoList {
			go getRepoLanguages(userName, repoName, repoChan)
		}
		for i := 0; i < len(repoList); i++ {
			tmpLanguageMap := <-repoChan
			for key, value := range tmpLanguageMap {
				checkData, exists := languageMap[key]
				if !exists {
					languageMap[key] = value
				} else {
					languageMap[key] = checkData + value
				}
			}
		}
		fmt.Println(languageMap)
		return c.File("hi World")
	})
	e.Logger.Fatal(e.Start(":8080"))
}

func getRepo(username string) ([]string, error) {
	client := github.NewClient(nil)
	ctx := context.Background()
	opt := &github.RepositoryListOptions{Type: "public"}
	repos, _, err := client.Repositories.List(ctx, username, opt)
	if err != nil {
		fmt.Println(err)
		return strings.Split("err err", " "), err
	}
	var reposList []string = make([]string, len(repos))
	repoChan := make(chan string)
	for _, repo := range repos {
		go getUrl(*repo.LanguagesURL, repoChan)
	}
	for i := 0; i < len(repos)-1; i++ {
		tmp := <-repoChan
		reposList[i] = tmp
	}
	return reposList, nil
}

func getUrl(url string, c chan<- string) {
	slicedURL := strings.Split(url, "/")
	c <- slicedURL[5]
}

func getRepoLanguages(userName string, repoName string, c chan<- map[string]int) {
	client := github.NewClient(nil)
	ctx := context.Background()
	languages, _, err := client.Repositories.ListLanguages(ctx, userName, repoName)
	if err != nil {
		fmt.Println(err)
	}
	c <- languages
	// fmt.Println(userName, repoName, languages)
}
