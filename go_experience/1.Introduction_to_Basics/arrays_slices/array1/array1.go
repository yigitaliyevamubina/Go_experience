//In the first step, 10 positive integers are supplied to standard input and must be written in order of input into a 10-element array. 
//The type of numbers included in the array must correspond to the smallest possible unsigned integer number. 
//The name of the array that you must create yourself is workArray (the condition is required). 
//To read from standard input, the fmt package is already imported.

// At the second stage, 3 more pairs of numbers are supplied to the standard input - the indices of the elements of this array that need to be swapped (if such a pair of numbers is 3 and 7, 
//then in the array the element with index 3 must be swapped with the element whose index is 7).

package main

import "fmt"

func main() {
	var a, x, y int
	var workArray [10]uint8
	var temp uint8

	for i := 0; i < 10; i++ {
	fmt.Scan(&a)
	workArray[i] = uint8(a)
	}

	for i := 0; i < 3; i++ {
	fmt.Scan(&x, &y)
	temp = workArray[x]
	workArray[x] = workArray[y]
	workArray[y] = temp
	}

	for i := 0; i < 10; i++ {
	fmt.Print(workArray[i], " ")
	}

}