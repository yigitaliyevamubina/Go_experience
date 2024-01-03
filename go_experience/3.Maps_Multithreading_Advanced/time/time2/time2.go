//The standard input is a string representation of the date and time of a specific event in the following format:
// 2020-05-15 08:00:00

// If the event time is before lunch (13-00), then you don’t need to change anything, just print the date to standard output in the same format.
// If the event should occur after lunch, you need to reschedule it for the same time the next day, and then output it to standard output in the same format.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	data, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic(err)
	}

	data = strings.TrimSuffix(data, "\n")
	parsedTime, err := time.Parse("2006-01-02 15:04:05", data)
	if err != nil {
		panic(err)
	}

	if parsedTime.Hour() >= 13 {
		parsedTime = parsedTime.AddDate(0, 0, 1)
	}
	fmt.Println(parsedTime.Format("2006-01-02 15:04:05"))
}