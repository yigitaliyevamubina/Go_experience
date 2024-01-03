//Ana two numbers. Determine the digits included in the entry of both the first and second numbers.
// The program should output the digits that are present in both numbers, separated by a space. 

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
		for k := 0; k < len(a); k++ {
			if string(a[k]) == string(n[i]) {
				ans += string(n[i])
				break
			}
		}
	}
	for i := 0; i < len(ans); i++ {
		fmt.Print(string(ans[i]), " ")
	}
}