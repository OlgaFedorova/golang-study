package main

import "fmt"

func main() {
	var num int8
	fmt.Scan(&num)

	if num == 1 {
		fmt.Printf("%d korova", num)
	} else if num >= 2 && num <= 4 {
		fmt.Printf("%d korovy", num)
	} else if num >= 5 && num <= 20 {
		fmt.Printf("%d korov", num)
	} else {
		if num%10 == 1 {
			fmt.Printf("%d korova", num)
		} else if num%10 == 2 || num%10 == 3 || num%10 == 4 {
			fmt.Printf("%d korovy", num)
		} else {
			fmt.Printf("%d korov", num)
		}
	}
}
