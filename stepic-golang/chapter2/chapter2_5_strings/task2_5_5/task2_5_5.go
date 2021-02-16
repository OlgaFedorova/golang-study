package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	for _, elem := range str {
		if strings.Count(str, string(elem)) > 1 {
			str = strings.ReplaceAll(str, string(elem), "")
		}
	}
	fmt.Println(str)

}
