package main

import "fmt"

func main() {
	var num uint32
	fmt.Scan(&num)
	var digRoot uint32
	digRoot = digitalRoot(num)
	for digRoot >= 10 {
		digRoot = digitalRoot(digRoot)
	}
	fmt.Println(digRoot)
}

func digitalRoot(num uint32) uint32 {
	var result uint32
	for num >= 10 {
		result = result + (num % 10)
		num = (num - (num % 10)) / 10
	}
	result = result + num
	return result
}
