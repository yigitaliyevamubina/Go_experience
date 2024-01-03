//You need to write a calculator function like this:
// func calculator(arguments <-chan int, done <-chan struct{}) <-chan int

// This function takes two read-only channels as an argument, returns a read-only channel.
// Through the arguments channel, the function will receive a series of numbers, and through the done channel, a signal about the need to complete work. When the completion signal is received, the function must send the sum of the received numbers to the output (return) channel.
// The calculator function must be non-blocking, returning control immediately.


package main

import "fmt"

func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
	sum := 0
	sm := make(chan int)
	go func() {
		for {
			select {
			case x, opened := <-arguments:
				if !opened {
					return
				}
				sum += x
			case <-done:
				sm <- sum
				close(sm)
				return
			}
		}
	}()
	return sm
}

func main() {
	arguments := make(chan int)
	done := make(chan struct{})

	go func() {
		for i := 0; i < 10; i++ {
			arguments <- i + 1
		}
		done <- struct{}{}
	}()

	channel := calculator(arguments, done)
	fmt.Println(<-channel)
}
