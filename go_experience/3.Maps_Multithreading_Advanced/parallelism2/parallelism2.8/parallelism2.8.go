package main

import "fmt"

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func() {
		for i := 0; i < n; i++ {
			n1 := <-in1
			n2 := <-in2
			go func() {
				n1 = fn(n1)
			}()

			go func() {
				n2 = fn(n2)
			}()
			out <- n1 + n2
		}
	}()
}

func fn(n int) int {
	return n
}

func main() {
	in1 := make(chan int, 10)
	in2 := make(chan int, 10)
	out := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			in1 <- i
			in2 <- i
		}
	}()

	merge2Channels(fn, in1, in2, out, 10)
	for i := 0; i < cap(out); i++ {
		fmt.Println(<-out)
	}
}
