//The standard input is a string representation of the date and time in the following format:
// 1986-04-16T05:20:00+06:00

// Your task is to convert this string to Time and then output it in UnixDate format:
// Wed Apr 16 05:20:00 +0600 1986

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	var t string
	reader, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic(err)
	}
	t = strings.TrimSuffix(reader, "\n")

	ans, err := time.Parse(time.RFC3339, t)
	if err != nil {
		log.Fatal(0)
	}
	fmt.Println(ans.Format(time.UnixDate))
}
