//The bank deposit is x rubles. Every year it increases by p percent, after which the fractional part of the kopecks is discarded. 
//Every year the deposit amount becomes larger. Determine how many years later the deposit will be at least y rubles.

package main 

import "fmt"

func main() {
	var x, p, y int 
	fmt.Scan(&x)
	fmt.Scan(&p)
	fmt.Scan(&y)
	cnt := 0
	for x < y {
		x += int((x * p) / 100)
		cnt++
	}
	fmt.Println(cnt)
}