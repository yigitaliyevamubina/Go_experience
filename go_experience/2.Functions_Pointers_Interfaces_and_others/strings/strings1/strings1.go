//The input is a string. You need to determine whether it is correct or not.
//A valid line begins with a capital letter and ends with a period. If the line is correct, print Right, otherwise, print Wrong.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.TrimRight(s, "\n")
    s2 := []rune(s)
    if unicode.IsUpper(s2[0]) && s2[len(s2)-1] == '.' {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}
