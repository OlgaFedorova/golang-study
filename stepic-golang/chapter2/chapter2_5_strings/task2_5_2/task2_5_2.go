package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scan(&str)
	text := []rune(str)
	isPalindrom := true
	for i := 0; i < len(text)/2; i++ {
		if text[i] != text[len(text)-i-1] {
			isPalindrom = false
			break
		}
	}

	if isPalindrom {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}
