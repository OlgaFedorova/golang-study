package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("stepic-golang/chapter3/chapter3_5_files/task3_5_3/task.data")
	if err == nil {
		rd := bufio.NewReader(file)
		var index int
		for str, err := rd.ReadString(';'); err != io.EOF || err == nil; str, _ = rd.ReadString(';') {
			index++
			if strings.Trim(str, ";") == "0" {
				break
			}
		}
		fmt.Println(index)
	}
	defer file.Close()

}
