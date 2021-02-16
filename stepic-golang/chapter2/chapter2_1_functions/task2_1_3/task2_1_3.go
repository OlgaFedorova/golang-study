package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	fmt.Println(vote(x, y, z))
}

func vote(x int, y int, z int) int {
	var count0, count1 int8
	if x == 0 {
		count0++
	} else {
		count1++
	}
	if y == 0 {
		count0++
	} else {
		count1++
	}
	if z == 0 {
		count0++
	} else {
		count1++
	}
	if count0 > count1 {
		return 0
	} else {
		return 1
	}
}
