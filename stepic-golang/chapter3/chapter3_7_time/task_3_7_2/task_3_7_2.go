package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	timeString, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	timeString = strings.Trim(timeString, "\n")

	layout := "2006-01-02 15:04:05"
	data, _ := time.Parse(layout, timeString)

	if data.Hour() < 13 {
		fmt.Println(data.Format(layout))
	} else {
		fmt.Println(data.AddDate(0, 0, 1).Format(layout))
	}
}
