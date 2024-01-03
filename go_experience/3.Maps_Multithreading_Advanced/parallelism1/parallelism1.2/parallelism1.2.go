//Write a function that takes a channel and a string. You need to send the same string to the specified channel 5 times, adding a space to it.

// The function should be called task2().

package main

import (
	"fmt"
	"sync"
)

func task2(channel chan string, s string) {
	for i := 0; i < 5; i++ {
		channel <- s + " "
	}
	close(channel)
}

func main() {
	channel := make(chan string, 5)
	var s string 
	fmt.Scan(&s)
	task2(channel, s)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if s, ok := <-channel; ok {
				fmt.Println(s)
			} else {
				return
			}
		}
	}()
	wg.Wait()
}
