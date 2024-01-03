//Given three natural numbers a, b, c. Determine whether there is a triangle with such sides.

package main

import "fmt"

func main() {
	var (
		a, b, c int
	)
	fmt.Scan(&a, &b, &c)
	if a + b > c && a + c > b && b + c > a{
		fmt.Println("Exists")
	} else {
		fmt.Println("Does not exist")
	}
}