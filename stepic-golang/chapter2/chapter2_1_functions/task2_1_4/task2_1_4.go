package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(fibonacci(n))
}

func fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	n1 := 1
	n2 := 1
	var res int
	for i := 3; i <= n; i++ {
		res = n1 + n2
		n1 = n2
		n2 = res
	}
	return res
}
