package main

import "fmt"

func main() {
	var num, n int
	fmt.Scan(&num, &n)
	nums := make([]int, 0, 10)
	for num >= 10 {
		nums = append(nums, num%10)
		num = (num - num%10) / 10
	}
	nums = append(nums, num)

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] != n {
			fmt.Print(nums[i])
		}
	}
}
