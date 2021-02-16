package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	fmt.Println(adding("%^80 ", "hhhhh20&&&&nd"))
}

func adding(a, b string) int64 {
	aAsNum, _ := convert(a)
	bAsNum, _ := convert(b)
	return aAsNum + bAsNum
}

func convert(a string) (int64, error) {
	result := make([]rune, 0, len(a))
	for _, symbol := range a {
		if unicode.IsDigit(symbol) {
			result = append(result, symbol)
		}
	}
	return strconv.ParseInt(string(result), 10, 64)
}
