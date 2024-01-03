//A three-digit number is given. Find the sum of its digits.

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	sum := n%10 + (n/10)%10 + n/100
	fmt.Println(sum)
}