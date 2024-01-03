//A three-digit number is given. Turn it over and then take it out.

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	sum := (n%10)*100 + ((n/10)%10)*10 + n/100
	fmt.Println(sum)
}