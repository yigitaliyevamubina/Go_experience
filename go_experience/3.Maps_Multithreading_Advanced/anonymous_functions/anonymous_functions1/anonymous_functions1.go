//you must declare a function of the form func(uint) uint, which can later be called by the name fn inside the main function.

// The function receives a positive integer number (uint) as input, returns a number of the same type in which there are no odd digits and the digit 0.
//If the resulting number is 0, then the number 100 is returned. The digits in the returned number have the same order as in the original number.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num uint
	fmt.Scan(&num)

	fn := func(num uint) uint {
		n := strconv.Itoa(int(num))
		nw := ""
		for _, v := range n {
			a, _ := strconv.Atoi(string(v))
			if a%2 == 0 && a > 0 {
				nw += string(v)
			}
		}
		if string(nw) == "" {
			return 100
		}
		ans, err := strconv.Atoi(nw)
		if err != nil {
			panic(err)
		}
		if ans == 0 {
			return 100
		} else {
			return uint(ans)
		}
	}

	fmt.Println(fn(num))
}
