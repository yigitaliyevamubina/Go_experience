//the first digit of the number

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for n >= 10 {
		n /= 10
	}
	fmt.Println(n)
}