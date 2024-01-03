//Remove a given digit from a natural number.
//Input data
//Enter the natural number and the digit to be delete.

package main

import (
	"fmt"
)

func main() {
	var n string
	fmt.Scan(&n)
	var a string 
	fmt.Scan(&a)
	ans := ""
	for i := 0; i < len(n); i++ {
		if string(n[i]) != a {
			ans += string(n[i])
		}
	}
	fmt.Println(ans)
}