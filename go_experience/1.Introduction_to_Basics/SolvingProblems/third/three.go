//The kth second of the day is passing. 
//Determine how many whole hours h and whole minutes m have passed since the beginning of the day. 
//For example, if k=13257=3*3600+40*60+57, then h=3 and m=40.

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println("It is", n/3600, "hours", (n%3600)/60, "minutes.")
}