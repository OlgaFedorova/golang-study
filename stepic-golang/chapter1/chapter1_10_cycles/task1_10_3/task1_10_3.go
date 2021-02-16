package main

import "fmt"

func main() {
	var n int
	var sum int
	fmt.Scan(&n)
	for ; n > 0; n-- {
		var a int
		fmt.Scan(&a)
		if a/10 > 0 && a/10 < 10 && a%8 == 0 {
			sum = sum + a
		}
	}
	fmt.Println(sum)
}
