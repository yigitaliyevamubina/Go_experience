package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n < 0 {
		fmt.Println("Число отрицательное")
	} else if n > 0 {
		fmt.Println("Число положительное")
	} else {
		fmt.Println("Ноль")
	}
}