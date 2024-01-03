//For standard input you receive 2 float numbers; as a result, you need to output the quotient of dividing the first number by the second accurate to four decimal places
//(in fact, after the dot, the result does not need to be converted to the original format).

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//example input: 1 149,6088607594936;1 179,0666666666666
	//example output: 0.9750

	s, err := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.Trim(s, "\n")
	if err != nil && err != io.EOF {
		log.Fatal(0)
	}
	s = strings.ReplaceAll(s, ",", ".")
	nums := strings.Split(s, ";")
	n1 := strings.ReplaceAll(nums[0], " ", "")
	n2 := strings.ReplaceAll(nums[1], " ", "")

	num1, err := strconv.ParseFloat(n1, 64)
	if err != nil {
		panic(err)
	}
	num2, err := strconv.ParseFloat(n2, 64)
	if err != nil {
		panic(err)
	}

	ans := num1 / num2
	fmt.Printf("%.4f", ans)
}