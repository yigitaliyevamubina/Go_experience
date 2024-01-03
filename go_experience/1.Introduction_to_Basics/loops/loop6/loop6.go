//Write a program that reads integers from the console, one number per line.

// For each number entered, check:

// if the number is less 10, then we skip this number;
// if the number is greater 100, then we stop reading the numbers;
// in other cases, print this number back to the console in a separate line.

package main 

import "fmt"

func main() {
	var n int 
	for {
		fmt.Scan(&n)
		if n > 100 {
			break
		} else if n < 10 {
			continue
		} else {
			fmt.Println(n)
		}
	}
}