package main

import "fmt"

func main() {
	var num float64
	fmt.Scan(&num)
	sq := num * num
	if num <= 0 {
		fmt.Printf("число %2.2f не подходит", num)
	} else if num > 10000 {
		fmt.Printf("%e", num)
	} else {
		fmt.Printf("%.4f", sq)
	}
}
