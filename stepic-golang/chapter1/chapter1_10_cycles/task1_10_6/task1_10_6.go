package main

import "fmt"

func main() {
	var n int
	numbers := make([]int, 0, 10)
	for fmt.Scan(&n); n <= 100; fmt.Scan(&n) {
		if n >= 10 {
			numbers = append(numbers, n)
		}
	}
	for _, elem := range numbers {
		fmt.Println(elem)
	}
}
