//Your task is to check whether the password entered by the user meets the specified requirements. 
//The password must be at least 5 characters long and must contain only Arabic numbers and Latin letters. 
//The input is provided with a password string. If the password meets the requirements, display "Ok", otherwise display "Wrong password".

package main

import (
	"fmt"
	"unicode"
)

func main() {
	var s string
	fmt.Scan(&s)
	s1 := []rune(s)
	if len(s1) >= 5 {
		check := true
		for i := 0; i < len(s1); i++ {
			if !unicode.Is(unicode.Latin, s1[i]) && !unicode.IsDigit(s1[i]) {
				check = false
				break
			}
		}
		if check {
			fmt.Println("Ok")
		} else {
				fmt.Println("Wrong password")
			}
		
	} else {
		fmt.Println("Wrong password")
	}
}