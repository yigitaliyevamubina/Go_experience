//The standard input takes 10 integers separated by spaces (numbers can be repeated).
//You need to calculate the result of executing the work function for each of the resulting numbers.
//The calculation results, separated by spaces, must be printed on a line.

// However, the work function takes too long. The execution time allocated to you is not enough to process each number sequentially, 
//so you need to implement caching of ready-made results and use them in your work.


package main

import "fmt"

func work(n int) int {
	if n >= 4 {
		return n + 1
	}
	return n - 1
}

func main() {
	var a int
	var ans = make(map[int]int, 10)
	for i := 0; i < 10; i++ {
		fmt.Scan(&a)
		if value, ok := ans[a]; ok {
			fmt.Print(value, " ")
		} else {
			ans[a] = work(a)
			fmt.Print(ans[a], " ")
		}
	}
}
