//Write a function called convert that converts an int64 input number to a uint16.

package main

import "fmt"

func convert(num int64) uint16 {
    return uint16(num)
}

func main() {
	var n int64
	n = -1	
	fmt.Println(convert(n))
}