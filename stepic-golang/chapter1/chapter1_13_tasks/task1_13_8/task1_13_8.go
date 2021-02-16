package main

import "fmt"

func main() {
	var n int
	var num, min int
	fmt.Scan(&n)
	var count uint8
	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		if i == 0 {
			min = num
		}

		if num == min {
			count++
		} else if num < min {
			min = num
			count = 1
		}
	}
	fmt.Println(count)
}
