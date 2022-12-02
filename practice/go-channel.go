package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	people := []string{"nico", "flynn", "vidigummy"}
	for _, person := range people {
		go isSexy(person, channel)
	}
	// result := <-channel
	fmt.Println("wating for messages")
	// 메세지가 채널에 들어온 만큼만 뽑아낼 수 있다. 초과할 시 deadlock 뜬다.
	for i := 0; i < len(people); i++ {
		fmt.Println("Receved this message : ", <-channel)
	}

	// fmt.Println("Receved this message : ", <-channel)
	//
	// fmt.Println("Receved this message : ", <-channel)
	// time.Sleep(time.Second * 10)
}

func isSexy(person string, c chan string) {
	// fmt.Println(person)
	time.Sleep(time.Second * 2)
	c <- person + "is sexy"
}
