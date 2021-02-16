package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)
	res := 1
	for res <= num {
		fmt.Print(res, " ")
		res = res * 2
	}
}
