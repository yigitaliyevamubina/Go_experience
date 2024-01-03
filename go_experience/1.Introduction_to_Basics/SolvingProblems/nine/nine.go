//Find the largest number on the segment from a to b that is a multiple of 7.

package main

import "fmt"

func main() {
	var a, b, check int
	fmt.Scan(&a, &b)
	check = 0
	for i := b; i >= a; i-- {
		if i % 7 == 0 {
			fmt.Println(i)
			check = 1
			break
		}
	} 
	if check == 0 {
		fmt.Println("NO")
	}
}