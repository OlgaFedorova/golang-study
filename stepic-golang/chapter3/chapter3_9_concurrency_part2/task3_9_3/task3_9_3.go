package main

import (
	"fmt"
)

func main() {
	firstChan := make(chan int)
	secondChan := make(chan int)
	stopChan := make(chan struct{})
	go func() {
		firstChan <- 5
	}()
	fmt.Println(<-calculator(firstChan, secondChan, stopChan))
}

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	done := make(chan int)
	go func() {
		defer close(done)
		select {
		case x := <-firstChan:
			done <- x * x
			return
		case x := <-secondChan:
			done <- 3 * x
			return
		case <-stopChan:
			return
		default:
			return
		}
	}()
	return done
}
