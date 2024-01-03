//The task is as follows: integer numbers in the range 0-100 are supplied to the standard input.
//Each number is supplied to the standard input from a new line (the number of numbers is not known).
//You need to read all these numbers and print their sum to standard output.

package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	buf := bufio.NewScanner(os.Stdin)
	var sum int
	for buf.Scan() {
		num, err := strconv.Atoi(buf.Text())
		if err != nil {
			return
		}
		sum += num
	}
	io.WriteString(os.Stdout, strconv.Itoa(sum))
}
