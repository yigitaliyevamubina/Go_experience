//example Go routines with Fibonacci

package main

import "fmt"

func fibonacci(channel chan int, quit chan int) {
	a, b := 0, 1
	for {
		select {
		case channel <- a:
			a, b = b, a+b
		case <-quit:
			fmt.Printf("%q\n", "Quit")
			return
		}
	}
}

func main() {
	channel := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-channel)
		}
		quit <- 0
	}()

	fibonacci(channel, quit)
}
