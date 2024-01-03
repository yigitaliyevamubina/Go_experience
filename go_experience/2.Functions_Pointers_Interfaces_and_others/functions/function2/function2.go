//Write a function that finds the smallest of the four numbers entered in the same function.

package main

import "fmt"

func minimumFromFour() int {
    var a int 
    minimum := 1000000000000 
    for i := 0; i < 4; i++ {
        fmt.Scan(&a)
        if a < minimum {
            minimum = a
        }
    }
    return minimum
}

func main() {
	fmt.Println(minimumFromFour())
}