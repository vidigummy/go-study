package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"pork", "beef", "chicken"}
	vidigummy := person{name: "vidigummy", age: 27, favFood: favFood}
	jibyo := person{"hyunjo", 24, favFood}
	fmt.Println(vidigummy.age - jibyo.age)
}
