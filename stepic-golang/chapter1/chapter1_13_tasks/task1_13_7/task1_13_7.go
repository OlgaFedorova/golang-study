package main

import "fmt"

func main() {
	var n int
	var num int
	fmt.Scan(&n)
	var count uint8
	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		if num == 0 {
			count++
		}
	}
	fmt.Println(count)
}
