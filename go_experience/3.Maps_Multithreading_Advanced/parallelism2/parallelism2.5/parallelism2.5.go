//example Timer with Ticker

package main

import (
	"fmt"
	"time"
)

func work() <-chan struct{} {
	quit := make(chan struct{})

	go func() {
		defer close(quit)

		stop := time.NewTimer(2*time.Second)
		
		tick := time.NewTicker(200*time.Millisecond)

		defer tick.Stop()

		for {
			select {
			case <- stop.C:
				return
			case <- tick.C:
				fmt.Println("Tic tac")
			}
		}

	}()
	return quit
}

func main() {
	<-work()
}
