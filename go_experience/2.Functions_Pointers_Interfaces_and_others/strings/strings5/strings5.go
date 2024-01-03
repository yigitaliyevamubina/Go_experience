//A string is given. You need to remove all characters that appear more than once and print the resulting string

package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	ans := ""
	for i := 0; i < len(s); i++ {
		if strings.Count(s, string(s[i])) <= 1 {
			ans += string(s[i])
		}
	}
	fmt.Println(ans)
}