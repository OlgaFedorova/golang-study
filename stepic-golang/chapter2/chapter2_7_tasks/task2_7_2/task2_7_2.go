package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func main() {
	var str string
	fmt.Scan(&str)

	var buffer bytes.Buffer
	for i, rune := range str {
		buffer.WriteRune(rune)
		if i != utf8.RuneCountInString(str)-1 {
			buffer.WriteRune('*')
		}
	}
	fmt.Println(buffer.String())
}
