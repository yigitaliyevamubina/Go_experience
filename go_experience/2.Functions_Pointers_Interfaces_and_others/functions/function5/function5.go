//Write a sumInt function that takes a variable number of int arguments and returns the number of arguments the function receives and their sum. 

package main

import "fmt"

func sumInt(nums ...int) (int, int){
    cnt := 0
    sum := 0
    for _, val := range nums {
        cnt++
        sum += val
    }
    return cnt, sum
}

func main() {
	a, b := sumInt(1, 0, 2, 4, 19) //example
	fmt.Println(a, b)
}