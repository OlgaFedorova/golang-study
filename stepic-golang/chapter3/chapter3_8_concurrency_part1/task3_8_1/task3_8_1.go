package main

import "fmt"

func main() {
	c := make(chan int, 1)
	task(c, 6)
	fmt.Println(<-c)
}

func task(c chan int, n int) {
	c <- n + 1
}
