//Using these numbers, determine the number of numbers that are equal to zero.

package main

import "fmt"

func main() {
	var n, a int
	fmt.Scan(&n)
	cnt := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		if a == 0 {
			cnt++
		}
	}
	fmt.Println(cnt)
}