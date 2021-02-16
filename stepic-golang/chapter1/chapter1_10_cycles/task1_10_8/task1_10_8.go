package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x)
	fmt.Scan(&y)

	n1 := nums(x)
	n2 := nums(y)

	for idx1 := len(n1) - 1; idx1 >= 0; idx1-- {
		for idx2 := len(n2) - 1; idx2 >= 0; idx2-- {
			if n1[idx1] == n2[idx2] {
				fmt.Print(n1[idx1], " ")
				break
			}
		}
	}
}

func nums(n int) []int {
	nums := make([]int, 0, 10)
	if n < 10 {
		nums = append(nums, n)
	} else {
		for n != 0 {
			nums = append(nums, n%10)
			n = (n - n%10) / 10
		}
	}
	return nums
}
