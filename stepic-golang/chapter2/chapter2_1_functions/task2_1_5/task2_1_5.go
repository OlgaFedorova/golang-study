package main

import "fmt"

func main() {
	a, b := sumInt(1, 0)
	fmt.Println(a, b)
}

func sumInt(s ...int) (int, int) {
	a := len(s)
	var b int
	for _, elem := range s {
		b = b + elem
	}
	return a, b
}
