package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4)
	fmt.Println("main() started")
	c := make(chan int, 3)
	done := make(chan struct{}, 1)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		c <- 4
		close(c)
		done <- struct{}{}
		close(done)
	}()
	fmt.Println(<-calculator(c, done))
	fmt.Println("main() stopped")
}

func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
		sum := 0
		cancel := false
		for !cancel {
			select {
			case x := <-arguments:
				sum = sum + x
			case <-done:
				cancel = true
			}
		}
		ch <- sum
		close(ch)
	}()
	return ch
}
