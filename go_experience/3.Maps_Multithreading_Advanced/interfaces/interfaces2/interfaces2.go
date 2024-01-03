//Your type should be such that when printed it will look like this: [ XXXX]: where the spaces are the "empty" capacity of the battery and the X is the "charged" capacity.
// Secondly, on standard input you receive a string consisting of exactly 10 digits: 0 or 1 (the order 0/1 is random).
//Your task is to read this string in any possible way and create, based on this string, an object of the type you declared in the first step: I hope you understand that the string symbolizes the capacity of the battery: 0 is the “empty” part, and 1 is the “charged” part.
// Third, the object you create must be called batteryForTest (using this name is mandatory).

package main

import (
	"fmt"
	"strings"
)

type Battery struct {
	bat string
}

func (b Battery) String() string {
	return "[" + strings.Repeat(" ", strings.Count(b.bat, "0")) + string(strings.Repeat("X", strings.Count(b.bat, "1"))) + "]"
}

func main() {
	var s string
	fmt.Scan(&s)
	batteryForTest := Battery{bat: s}
	fmt.Print(batteryForTest)
}
