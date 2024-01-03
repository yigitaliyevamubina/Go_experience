//An integer is given as input. It is necessary to square each digit of the number and print the resulting number.

package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	for i := 0; i < len(s); i++ {
		n := s[i] - '0'
		fmt.Print(n * n)		
	}
}