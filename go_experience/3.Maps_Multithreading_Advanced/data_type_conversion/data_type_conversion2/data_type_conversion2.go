//Imagine that you work in a large company that uses modular architecture.
//Your colleague wrote a module with some logic (you don’t know) and passes some data to your program.
//You are writing a function that reads two variables of type string and returns a number of type int64 that needs to be obtained by adding these strings.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func adding(s1 string, s2 string) int64 {
	a := []rune{}
	b := []rune{}
	for _, num := range s1 {
		if unicode.IsDigit(num) {
			a = append(a, num)
		}
	}
	for _, num := range s2 {
		if unicode.IsDigit(num) {
			b = append(b, num)
		}
	}
	num1, err := strconv.Atoi(string(a))
	if err != nil {
		return -1
	}
	num2, err := strconv.Atoi(string(b))
	if err != nil {
		return -1
	}
	return int64(num1 + num2)
}

func main() {
	//example input: %^80 hhhhh20&&&&nd
	//example output: 100 (explanation: 80 + 20)
	var s []string
	reader, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.Split(reader, " ")
	fmt.Println(adding(s[0], s[1]))
}
