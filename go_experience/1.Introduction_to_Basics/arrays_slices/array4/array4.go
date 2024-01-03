//Given a sequence consisting of integers. 
//Write a program that counts the number of positive numbers among the elements of a sequence.

package main

import "fmt"

func main() {
	var a, n int
	fmt.Scan(&n)
	lst := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		lst[i] = a
	}
	cnt := 0
	for i := 0; i < n; i++ {
		if lst[i] >= 0 {
			cnt++
		}
	}
	fmt.Println(cnt)

}