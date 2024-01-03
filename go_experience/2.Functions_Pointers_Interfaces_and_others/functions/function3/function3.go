//Write a “voting function” that returns the value (0 or 1) that occurs most often among the values ​​of its arguments x, y, z.

package main

import "fmt"

func vote(x int, y int, z int) int {
    if x == y {
        return x
    }
    return z
}

func main() {
	var a, b, c int 
	fmt.Scan(&a, &b, &c)
	fmt.Println(vote(a, b, c))
}