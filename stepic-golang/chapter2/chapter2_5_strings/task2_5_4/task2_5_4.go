package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
	strAsRune := []rune(str)
	for i := 0; i < len(strAsRune); i++ {
		strAsRune = append(strAsRune[:i], strAsRune[i+1:]...)
	}
	fmt.Println(string(strAsRune))
}
