package main

import "fmt"

func main() {
	var a, n int
	fmt.Scan(&a)
	if a == 1 {
		fmt.Println(1)
		return
	}

	n1 := 1
	n2 := 1
	idx := 2
	for n < a {
		n = n1 + n2
		n1 = n2
		n2 = n
		idx++
	}

	if a == n {
		fmt.Println(idx)
	} else {
		fmt.Println(-1)
	}
}
