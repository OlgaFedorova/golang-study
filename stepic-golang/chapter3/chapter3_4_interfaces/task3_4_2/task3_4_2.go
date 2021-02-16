package main

import (
	"fmt"
	"strings"
)

type Battery struct {
	full uint
}

func (a Battery) String() string {
	return fmt.Sprintf("[%10s]", strings.Repeat("X", int(a.full)))
}

func main() {
	var text string
	fmt.Scan(&text)
	var fullCount uint
	for _, r := range text {
		if int(r-'0') == 1 {
			fullCount++
		}
	}
	batteryForTest := Battery{
		full: fullCount,
	}

	fmt.Println(batteryForTest)
}
