package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/google/go-github/github"
	"github.com/labstack/echo"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

type userInfoGithub struct {
	login     string
	id        int
	repos_url string
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		testGithubApi("vidigummy")
		return c.File("home.html")
	})
	e.Logger.Fatal(e.Start(":8080"))
}

func testGithubApi(username string) error {
	client := github.NewClient(nil)
	ctx := context.Background()
	// opt := &github.RepositoryListByOrgOptions{Type: "public"}
	opt := &github.RepositoryListOptions{Type: "public"}
	// repos, _, err := client.Repositories.ListByOrg(ctx, username, opt)
	repos, _, err := client.Repositories.List(ctx, username, opt)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(repos)
	for _, repo := range repos {
		fmt.Println(repo)
	}
	return nil
}

func getUserInfo(username string) (string, error) {
	userApiURL := "https://api.github.com/users/" + username
	// vidigummyInfo := new(userInfoGithub)
	// err := getJson(userApiURL, vidigummyInfo)
	err := getUserInfoFromURL(userApiURL)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(vidigummyInfo.repos_url)
	return "OK", nil
}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func getUserInfoFromURL(url string) error {
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	var data userInfoGithub
	json.Unmarshal(body, &data)
	fmt.Printf("Results: %v\n", data)
	return nil
}
