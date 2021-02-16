package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	n1 := n % 10
	n2 := ((n - n1) / 10) % 10
	n3 := ((n-n1)/10 - n2) / 10
	switch {
	case n1 != n2 && n1 != n3 && n2 != n3:
		fmt.Println("YES")
	default:
		fmt.Println("NO")
	}
}
