package main

import "fmt"

func main() {
	c := make(chan string, 6)
	task2(c, "str6")
	fmt.Println(<-c)
}

func task2(c chan string, n string) {
	c <- n + " "
	c <- n + " "
	c <- n + " "
	c <- n + " "
	c <- n + " "
}
