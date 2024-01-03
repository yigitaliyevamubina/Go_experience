//Write a program that, in a sequence of numbers, finds the sum of two-digit numbers that are multiples of 8. 
//The program in the first line receives as input the number n - the number of numbers in the sequence, 
//in the second line - n numbers included in this sequence.

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	lst := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&lst[i])
	}
	sum := 0
	for _, value := range lst {
		if value%8 == 0 && value < 100 && value > 9{
			sum += value
		}
	}
	fmt.Println(sum)
}