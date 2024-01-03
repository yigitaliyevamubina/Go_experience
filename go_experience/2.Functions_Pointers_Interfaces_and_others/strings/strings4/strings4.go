//A string is given as input, you need to make another string from it, leaving only odd characters (counting from zero).


package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	ans := ""
	for i := 0; i < len(s); i++ {
		if i % 2 == 1 {
			ans += string(s[i])
		}
	}
	fmt.Println(ans)
}