package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			work()
		}()
	}
	wg.Wait()
}

func work() {
	fmt.Println("Горутина начала выполнение")
	time.Sleep(2 * time.Second)
	fmt.Println("Горутина завершила выполнение")
}
