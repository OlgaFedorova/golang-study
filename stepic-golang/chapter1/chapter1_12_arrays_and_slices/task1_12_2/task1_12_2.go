package main

import "fmt"

func main() {
	var sl = make([]int, 0, 10)
	var n, in int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&in)
		sl = append(sl, in)
	}
	fmt.Println(sl[3])
}
