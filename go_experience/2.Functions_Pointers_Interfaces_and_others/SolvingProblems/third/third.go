//Given a string containing only Arabic numerals. Find and display the largest number.

package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	maxi := ""
	for i := 0; i < len(s); i++ {
		if string(s[i]) > maxi {
			maxi = string(s[i])
		}
	}
	fmt.Println(maxi)
}
