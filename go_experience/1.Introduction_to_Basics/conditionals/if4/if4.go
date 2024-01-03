//checking if the sum of first three digits of the six-digit number from the left equals to the sum of last three digits

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	first := 0
	last := 0
	for n >= 1000 {
		first += n % 10
		n /= 10
	}
	for n > 0 {
		last += n % 10
		n /= 10
	}
	if first == last {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}