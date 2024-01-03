//The input is five integers, which are written to an array.
// You need to write a piece of code that can be used to find and print the maximum number in this array.

package main

import "fmt"

func maxi(lst []int) int {
	maximum := lst[0]
	for _, value := range lst {
		if value > maximum {
			maximum = value
		}
	}
	return maximum
}

func main() {
	lst2 := make([]int, 5)
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		lst2[i] = a
	}
	fmt.Println(maxi(lst2))
}