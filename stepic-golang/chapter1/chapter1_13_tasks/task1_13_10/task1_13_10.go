package main

import (
	"fmt"
)

func main() {
	var a, b, step int
	fmt.Scan(&a, &b)
	if a < b {
		step = -1
	} else {
		step = 1
	}

	for i := b; i >= a; i = i + step {
		if i%7 == 0 {
			fmt.Println(i)
			return
		}
	}
	fmt.Println("NO")
}
