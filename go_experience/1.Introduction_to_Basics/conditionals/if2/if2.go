//given three-digit number
//print 'YES' if it has only distinct digits, else 'NO'

package main

import "fmt"

func main() {
	var n string
	fmt.Scanln(&n)

	if int(n[0]) == int(n[1]) && int(n[1]) == int(n[2]) {
		fmt.Println("NO")
    } else if int(n[0]) == int(n[1]) || int(n[1]) == int(n[2]) || int(n[0]) == int(n[2]) {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}