//The Fibonacci sequence is defined as follows: φ1=1, φ2=1, φn=φn-1+φn-2 for n>1. The beginning of the Fibonacci series looks like this: 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, ... 
//Write a function that, given a natural number n, returns φn.

package main

import "fmt"

func fibonacci(n int) int {
    indx := 2
    a, b := 0, 1 
    for indx <= n {
        indx++ 
        a, b = b, a + b 
    } 
    return b 
}

func main() {
	var n int 
	fmt.Scan(&n)
	fmt.Println(fibonacci(n))
}