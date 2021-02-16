package main

import "fmt"

func main() {
	var k uint32
	fmt.Scan(&k)
	sec := k % 60
	min := ((k - sec) / 60) % 60
	h := (k - min - sec) / 3600
	fmt.Printf("It is %d hours %d minutes.", h, min)
}
