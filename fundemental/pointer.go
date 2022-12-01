package main

import "fmt"

func main() {
	// &변수명 하면 포인터를 리턴한다.
	// *변수명 하면 그 포인터 주소에 있는 값을 가져온다.
	a := 2
	b := &a
	*b = 20
	fmt.Println(a, *b)
}
