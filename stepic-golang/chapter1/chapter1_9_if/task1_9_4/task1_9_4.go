package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	n1 := n
	var nums [6]int
	idx := 5
	for n1 >= 10 {
		nums[idx] = n1 % 10
		n1 = (n1 - n1%10) / 10
		idx--
	}
	nums[0] = n1

	if nums[0]+nums[1]+nums[2] == nums[3]+nums[4]+nums[5] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
