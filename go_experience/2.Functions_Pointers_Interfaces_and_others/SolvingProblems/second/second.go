//Given a string containing only English letters (upper and lowercase). 
//Add the symbol ‘*’ (asterisk) between letters (there is no need to add the symbol ‘*’ before the first letter and after the last).

package main

import (
	"fmt"
	"unicode"
)

func main() {
	var s string
	fmt.Scan(&s)
	answer := ""
	for i := 0; i < len(s)-1; i++ {
		if unicode.IsLetter(rune(s[i])) {
			answer += string(s[i]) + "*"
		}
	}
	answer += string(s[len(s)-1])
	fmt.Println(answer)
}