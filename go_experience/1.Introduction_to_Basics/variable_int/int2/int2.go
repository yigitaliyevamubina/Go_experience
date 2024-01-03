package main

import "fmt"

func main() {

	var a int
	var b int
	var c int

	fmt.Scan(&a) // input 'a' variable
	fmt.Scan(&b) // input 'b' variable

	a *= a
	b *= b
	c = a + b
	fmt.Println(c)
}
