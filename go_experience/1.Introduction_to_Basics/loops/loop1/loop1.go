//Write a program that prints the squares of natural numbers from 1 to 10. 
//The square of each number should be printed on a new line. 

package main

import "fmt"

func main() {
	for i := 1; i < 11; i++ {
		fmt.Println(i * i)
	}

}