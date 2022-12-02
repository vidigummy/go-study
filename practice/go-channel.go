package main

import (
	"fmt"
	"time"
)

func main() {
	// 채널로 뭘 받을지 정확히 해야한다.
	// 고루틴이고 뭐고 wait 안 하면 함수는 끝나고 고아 쓰레드는 버려짐. 그러니까 받기로 했으면 wating 시켜줘야함
	// 채널을 사용하는 방법은 화살표다. channel <- 변수 는 채널에다 변수 값 집어 넣는거고, 변수 <- channel은 변수에다 채널 맨 앞에 있는거 하나씩 집어넣는거다.(근데 끝나는 순서 보장은 안 되는듯)
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
