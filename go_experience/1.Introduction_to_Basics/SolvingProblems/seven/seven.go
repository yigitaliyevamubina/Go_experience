//Find the number of minimal elements in the sequence.

package main

import "fmt"

func min(lst []int, n int) int {
	mini := lst[0]
	for i := 1; i < n; i++ {
		if lst[i] < mini {
			mini = lst[i]
		}
	}
	return mini
}

func main() {
	var n, a int
	fmt.Scan(&n)
	lst := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		lst[i] = a
	}
	elem := min(lst, n)
	cnt := 0
	for i := 0; i < n; i++ {
		if lst[i] == elem {
			cnt++
		}
	}
	fmt.Println(cnt)
}