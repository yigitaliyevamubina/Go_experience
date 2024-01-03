//Given a natural number A > 1. Determine what Fibonacci number it is, that is, print a number n such that φn=A. 
//If A is not a Fibonacci number, print the number -1.

package main

import "fmt"

func isFibonacci(n int) int {
	if n < 0 {
		return  -1
	}
	a, b := 0, 1 
	indx := 1
	for b <= n {
		if b == n {
			return indx
		}
		a, b = b, a + b 
		indx++
	}
	return -1
}

func main() {
	var num int 
	fmt.Scan(&num)
	fmt.Println(isFibonacci(num))
}