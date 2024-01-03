//The sequence consists of natural numbers and ends with the number 0. 
//Determine the number of elements of this sequence that are equal to its largest element.

package main

import (
	"fmt"
)

func max(lst []int) int {
	val := lst[0]
	for i := 1; i < len(lst); i++ {
		if lst[i] > val {
			val = lst[i]
		}
	}
	return val
}

func main() {
	var n int
	lst := []int{}
	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		lst = append(lst, n)
	}
	max_val := max(lst)
	cnt := 0
	for _, value := range lst {
		if value == max_val {
			cnt++
		}
	}
	fmt.Println(cnt)
}
