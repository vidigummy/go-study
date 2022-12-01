package main

func canIDrink(age int) bool {
	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 18:
		return true
	}
	return true
}

func main() {
}
