//Given a number N, print all integer values ​​of powers of two not exceeding N, in ascending order.

package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := 1
	k := 0
	for a <= n {
		fmt.Print(a, " ")
		k++
		a = int(math.Pow(2, float64(k)))
	}
}