package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const now = 1589570165

func main() {
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	input = strings.Trim(input, "\n")
	input = strings.Replace(input, " мин. ", "m", 1)
	input = strings.Replace(input, " сек.", "s", 1)
	dur, _ := time.ParseDuration(input)

	fmt.Println(time.Unix(now, 0).Add(dur).Format(time.UnixDate))
}
