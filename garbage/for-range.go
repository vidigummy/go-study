package main

import "fmt"

func superAdd(numbers ...int) int {
	fmt.Println(numbers)
	for number := range numbers {
		fmt.Println(number)
	}
	return 1
}

func superAddd(numbers ...int) int {
	fmt.Println(numbers)
	for number := range numbers {
		fmt.Println(numbers[number])
	}
	return 1
}

func superAddWithIndex(numbers ...int) int {
	result := 0
	for index, number := range numbers {
		fmt.Println(index, number)
		result += number
	}
	return result
}

func main() {
	result := superAddWithIndex(1, 2, 3, 45, 5234)
	fmt.Println(result)
}
