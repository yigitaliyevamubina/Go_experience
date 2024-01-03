//You need to write a calculator function like this:
// func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int

// The function takes 3 channels as arguments and returns a channel of type <-chan int.

// in case the argument is received from the firstChan channel, you must send the square of the argument to the output (return) channel.
// in case the argument is received from the secondChan channel, you must send the result of multiplying the argument by 3 to the output (return) channel.
// in case the argument is received from the stopChan channel, you just need to terminate the function.

// The calculator function must be non-blocking, returning control immediately. Your function will receive only one value into one of the channels - they received the value, processed it, and completed their work.

package main

import (
	"fmt"
	"sync"
)

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	channel := make(chan int)
	go func(channel chan int) {
		defer close(channel)
		select {
		case num := <-firstChan:
			channel <- num * num
		case num := <-secondChan:
			channel <- num * 3
		case <-stopChan:
			return
		}
	}(channel)
	return channel
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	firstChan := make(chan int)
	secondChan := make(chan int)
	stopChan := make(chan struct{})
	go func() {
		defer wg.Done()
		firstChan <- 10
		secondChan <- 15
		stopChan <- struct{}{}
	}()

	go func() {
		defer close(firstChan)
		defer close(secondChan)
		defer close(stopChan)
		wg.Wait()
	}()

	fmt.Println(<-calculator(firstChan, secondChan, stopChan))

}
