package main

import "fmt"

func main() {
	var a, b uint32
	fmt.Scan(&a, &b)
	var mid float32
	mid = (float32(a) + float32(b)) / 2.0
	fmt.Println(mid)
}
