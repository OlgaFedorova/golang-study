package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	res := 0
	for scanner.Scan() {
		str := scanner.Text()
		strAsInt, err := strconv.Atoi(str)
		if err == nil {
			res = res + strAsInt
		}
	}
	io.WriteString(os.Stdout, strconv.Itoa(res))

}
