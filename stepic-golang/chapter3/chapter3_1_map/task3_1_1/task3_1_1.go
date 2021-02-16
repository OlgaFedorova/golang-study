package main

import "fmt"

func main() {
	var nums [10]int
	for i := 0; i < 10; i++ {
		fmt.Scan(&nums[i])
	}
	cache := make(map[int]int)
	for i := 0; i < 10; i++ {
		if value, inCache := cache[nums[i]]; inCache {
			fmt.Print(value, " ")
		} else {
			value := work(nums[i])
			cache[nums[i]] = value
			fmt.Print(value, " ")
		}
	}
}

func work(x int) int {
	return x + 6
}
