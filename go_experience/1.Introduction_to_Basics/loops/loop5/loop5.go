//Find  the first  number from 1 to n inclusive that is a multiple of c but NOT a multiple of d.

package main 

import "fmt"

func main() {
	var n, c, d int 
	fmt.Scan(&n)
	fmt.Scan(&c)
	fmt.Scan(&d)

	for i := 1; i <= n; i++ {
		if i % c == 0 && i % d != 0 {
			fmt.Println(i)
			break
		}
	}
}