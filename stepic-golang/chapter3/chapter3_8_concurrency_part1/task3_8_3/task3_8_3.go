package main

import "fmt"

func main() {
	inputStream := make(chan string, 6)
	inputStream <- "str1"
	inputStream <- "str2"
	inputStream <- "str1"
	inputStream <- "str1"
	inputStream <- "str2"
	inputStream <- "str3"
	close(inputStream)
	outputStream := make(chan string, 6)
	removeDuplicates(inputStream, outputStream)

	for v := range outputStream {
		fmt.Println(v)
	}
}

func removeDuplicates(inputStream, outputStream chan string) {
	previous := <-inputStream
	outputStream <- previous
	for v := range inputStream {
		if v != previous {
			previous = v
			outputStream <- previous
		}
	}
	close(outputStream)
}
