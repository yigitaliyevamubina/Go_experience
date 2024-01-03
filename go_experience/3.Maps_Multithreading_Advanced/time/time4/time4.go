//The standard input provides data on the duration of the period (the format is shown in the example). Additionally, you are given a date in Unix-Time format: 1589570165 as an int64 constant (nanoseconds are 0 for conversion purposes).

// You need to read data about the duration of a period, convert it to the Duration type, and then output (in UnixDate format) the date and time resulting from adding the period to a standard date.

package main

import (
	"fmt"
	"time"
)

func main() {
	const now = 1589570165
	var (
		min int
		sec int
	)
	fmt.Scanf("%d мин. %d сек.", &min, &sec)
	seconds := 60*min + sec
	seconds += now
	
	t := time.Unix(int64(seconds), 0).UTC()

	fmt.Println(t.Format(time.UnixDate))
}
