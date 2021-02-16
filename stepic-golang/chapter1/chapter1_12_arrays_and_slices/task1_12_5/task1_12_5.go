package main

import "fmt"

func main() {
	var n, in int
	fmt.Scan(&n)
	arr := make([]int, 0, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&in)
		arr = append(arr, in)
	}
	var count int
	for _, elem := range arr {
		if elem > 0 {
			count++
		}
	}
	fmt.Println(count)
}
