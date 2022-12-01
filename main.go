package main

import (
	"fmt"
	"strings"
)

// naked return
func lenAndUpper(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

// defer
// excute somthing when function is done
func lLenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done!")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}
func main() {
	length, uppercase := lLenAndUpper("vidigummy")
	fmt.Println(length, uppercase)
}
