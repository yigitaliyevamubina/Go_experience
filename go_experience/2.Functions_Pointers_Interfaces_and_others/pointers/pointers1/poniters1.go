//Write a function that multiplies the values ​​referenced by two pointers and prints the resulting product to the console.

package main

import "fmt"

func test(x1 *int, x2 *int) {
    fmt.Println(*x1 * *x2)
}

func main() {
	var a, b int 
	fmt.Scan(&a, &b)
	test(&a, &b)
}