package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	var psw string
	fmt.Scan(&psw)
	if utf8.RuneCountInString(psw) >= 5 {
		isWrong := false
		for _, elem := range psw {
			if !unicode.Is(unicode.Latin, elem) && !unicode.Is(unicode.Number, elem) {
				isWrong = true
				break
			}
		}
		if isWrong {
			fmt.Println("Wrong password")
		} else {
			fmt.Println("Ok")
		}

	} else {
		fmt.Println("Wrong password")
	}

}
