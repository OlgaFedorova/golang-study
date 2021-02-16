package main

import (
	"fmt"
)

func main() {
	var str string
	var max int32
	fmt.Scan(&str)

	for i, rune := range str {
		if i == 0 {
			max = rune
		} else if rune > max {
			max = rune
		}
	}
	fmt.Println(string(max))
}
