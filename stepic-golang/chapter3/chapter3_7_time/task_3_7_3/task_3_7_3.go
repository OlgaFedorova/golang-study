package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	input = strings.Trim(input, "\n")
	timesString := strings.Split(input, ",")

	layout := "02.01.2006 15:04:05"
	first, _ := time.Parse(layout, timesString[0])
	second, _ := time.Parse(layout, timesString[1])

	if first.After(second) {
		fmt.Println(first.Sub(second))
	} else {
		fmt.Println(second.Sub(first))
	}
}
