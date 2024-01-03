//Given an array consisting of integers. The numbering of elements starts from 0. 
//Write a program that will print the elements of an array whose indices are even (0, 2, 4...).

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
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(lst[i], " ")
		}
	}

}