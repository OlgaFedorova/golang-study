package main

import "fmt"

func main() {
	var n int
	arr := make([]int, 0, 10)
	for fmt.Scan(&n); n != 0; {
		arr = append(arr, n)
		fmt.Scan(&n)
	}
	max := arr[0]
	for _, elem := range arr {
		if elem > max {
			max = elem
		}
	}
	var count int
	for _, elem := range arr {
		if elem == max {
			count++
		}
	}
	fmt.Println(count)
}
