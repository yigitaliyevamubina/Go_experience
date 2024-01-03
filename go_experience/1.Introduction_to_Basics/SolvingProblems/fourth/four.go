//Three numbers are given - a,b,c (a<b<c) - lengths of the sides of the triangle. 
//We need to check whether the triangle is right-angled. If it is, print "Rectangular". Otherwise output "Non-rectangular".

package main

import "fmt"

func main() {
	var (
		a, b, c int
	)
	fmt.Scan(&a, &b, &c)
	if a*a+b*b == c * c{
		fmt.Println("Rectangular")
	} else {
		fmt.Println("Non-rectangular")
	}
}