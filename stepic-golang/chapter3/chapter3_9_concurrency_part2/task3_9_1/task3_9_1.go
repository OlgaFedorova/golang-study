package main

import "fmt"

func main() {
	done := make(chan struct{})
	go func() {
		defer close(done)
		work()
	}()
	<-done
}

func work() {
	fmt.Println("Do something")
}
