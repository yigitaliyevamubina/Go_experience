//The input is a string. We need to determine whether it is a palindrome. 
//If the string is a palindrome, print Palindrome, otherwise, print No. All input lines are in lowercase.

package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	txt := []rune(s)
	check := true
	for i, k := 0, len(txt)-1; i < k; i, k = i+1, k-1 {
		if txt[i] != txt[k] {
			check = false
			break
		}
	}
	if check {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}