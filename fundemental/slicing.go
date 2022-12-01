package main

import "fmt"

func main() {
	names := []string{"nico", "vidi", "jibyo", "fififi", "foo"}
	names[3] = "fjijiji"
	names = append(names, "vidigummy")
	fmt.Println(names)
}
