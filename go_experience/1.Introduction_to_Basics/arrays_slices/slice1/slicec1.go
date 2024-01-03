//Write a program that takes a number as input N (N≥4), and then N integer slice elements. 
//The output should be the 4th (3rd by index) element of this slice.

package main

import "fmt"

func main() {

	var n, a int
	fmt.Scan(&n)
	var lst = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		lst[i] = a
	}
	fmt.Println(lst[3])

}