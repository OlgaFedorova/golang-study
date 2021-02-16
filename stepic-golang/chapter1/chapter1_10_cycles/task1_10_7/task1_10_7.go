package main

import "fmt"

func main() {
	var x, p, y int
	fmt.Scan(&x)
	fmt.Scan(&p)
	fmt.Scan(&y)
	res := x
	years := 0
	for res < y {
		res = res + (res * p / 100)
		years++
	}
	fmt.Println(years)
}
