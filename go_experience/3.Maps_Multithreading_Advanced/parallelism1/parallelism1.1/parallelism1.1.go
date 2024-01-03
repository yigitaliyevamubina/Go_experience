//Write a function that takes a channel and a number N of type int. It is necessary to return the value N+1 to the channel.
// The function should be called task().

package main

import "fmt"

func task(channel chan int, n int) {
	channel <- n + 1
}

func main() {
	channel := make(chan int, 1)
	task(channel, 10)
	fmt.Println(<-channel)
}
