//The standard input is a string representation of two dates separated by a comma (see example for data format).

// You need to convert the received data to the Time type, and then display the duration of the period between the smaller and larger dates.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	data, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}

	data = strings.TrimSuffix(data, "\n")
	t := strings.Split(data, ",")

	parsed1, _ := time.Parse("02.01.2006 15:04:05", t[0])
	parsed2, _ := time.Parse("02.01.2006 15:04:05", t[1])

	var ans time.Duration
	if parsed1.After(parsed2) {
		ans = time.Since(parsed2) - time.Since(parsed1)
	} else {
		ans = time.Since(parsed1) - time.Since(parsed2)
	}

	fmt.Println(ans.Round(time.Second))

	// fmt.Println(time.Since(parsed1))
	// fmt.Println(time.Since(parsed2))
}
