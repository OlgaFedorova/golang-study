package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	f(s)
}

func f(text string) {
	fmt.Println(text)
}
