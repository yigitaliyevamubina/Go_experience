//You need to call the work() function 10 times in separate goroutines and wait for the results of the called functions.

package main

import (
	"fmt"
	"sync"
	"time"
)

func work() {
	fmt.Println("Goroutine starts")
	time.Sleep(5 * time.Second)
	fmt.Println("Goroutine stops")
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			work()
		}()
	}
	wg.Wait()
}
