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
	for idx := range arr {
		if idx%2 == 0 {
			fmt.Print(arr[idx], " ")
		}
	}
}
