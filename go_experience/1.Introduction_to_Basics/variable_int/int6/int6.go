package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
    fmt.Print("It is ", a/30, " hours ", (a%30) * 2, " minutes.")
}