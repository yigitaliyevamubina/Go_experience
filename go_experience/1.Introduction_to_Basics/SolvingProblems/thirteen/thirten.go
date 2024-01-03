//Given a natural number N. Print its representation in binary form.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	var ans string
	for n > 0 {
		ans = strconv.Itoa(n % 2) + ans
		n /= 2
	}
	fmt.Println(ans)
}