package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.Trim(text, "\n")
	nums := strings.Split(text, ";")
	a, _ := convert(nums[0])
	b, _ := convert(nums[1])
	fmt.Printf("%.4f", a/b)

}

func convert(a string) (float64, error) {
	a = strings.ReplaceAll(a, " ", "")
	a = strings.ReplaceAll(a, ",", ".")
	return strconv.ParseFloat(a, 64)
}
