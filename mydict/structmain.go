package main

import (
	"fmt"

	"example.com/m/v2/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseword := "hello"
	dictionary.Add(baseword, "greetings")
	word, _ := dictionary.Search(baseword)
	fmt.Println("삭제 전 : ", word)
	dictionary.Delete(baseword)
	deltedWord, err := dictionary.Search(baseword)
	if err == nil {
		fmt.Println("성공 : ", deltedWord)
	} else {
		fmt.Println(err)
	}
	delErr := dictionary.Delete(baseword)
	if delErr != nil {
		fmt.Println(delErr)
	} else {
		fmt.Println("왜 되냐")
	}
}
