package main

import (
	"fmt"
	"time"
)

func main() {
	var timeString string
	fmt.Scan(&timeString)
	data, _ := time.Parse(time.RFC3339, timeString)
	fmt.Println(data.Format(time.UnixDate))
}
