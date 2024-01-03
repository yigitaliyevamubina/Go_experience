//Given two strings X and S. You need to find and print the first occurrence of the substring S in string X. 
//If the substring S is not in string X, print -1.

package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1 string
	fmt.Scan(&s1)
	var s2 string
	fmt.Scan(&s2)
	fmt.Println(strings.Index(s1, s2))
}