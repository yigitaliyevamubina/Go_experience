//The digital root of a natural number is a digit obtained as a result of an iterative process of summing digits, at each iteration of which the result obtained at the previous iteration is taken to calculate the sum of digits. 
//This process is repeated until one digit is obtained.

// For example, the digital root 65536 is 7, because 6+5+5+3+6=25 and 2+5=7.

// Given a number, determine its digital root.

package main 

import "fmt"

func main() {
	var n int 
	fmt.Scan(&n)
	for n > 10 {
		sm := 0
		for n > 0 {
			sm += n % 10 
			n /= 10 
		}
		n = sm 
	}
	fmt.Println(n)
}