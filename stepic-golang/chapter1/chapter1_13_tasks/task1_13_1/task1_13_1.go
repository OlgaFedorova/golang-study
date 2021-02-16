package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	num1 := num % 10
	num2 := ((num - num1) / 10) % 10
	num3 := (num - num2 - num1) / 10 / 10

	result := num1 + num2 + num3
	fmt.Println(result)
}
