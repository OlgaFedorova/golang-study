package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	n1 := n
	for n1 >= 10 {
		n1 = (n1 - n1%10) / 10
	}
	fmt.Println(n1)
}
