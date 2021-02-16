package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)

	for _, r := range str {
		res := int(r-'0') * int(r-'0')
		fmt.Print(res)
	}
}
