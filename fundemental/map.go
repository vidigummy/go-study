package main

import "fmt"

func main() {
	vidi := map[string]string{"name": "vidi", "age": "27"}
	for key, _ := range vidi {
		fmt.Println(key)
	}
}
