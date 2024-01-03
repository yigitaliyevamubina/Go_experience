//Write a function f() that will accept the string text and print it on the screen.

package main

import "fmt"

func f(text string) {
    fmt.Println(text)
}

func main() {
	var s string 
	fmt.Scan(&s)
	f(s)
}

