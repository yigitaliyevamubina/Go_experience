//Swap the values ​​of the variables referenced by the pointers. After this, the variables need to be output.

package main

import "fmt"

func test(x1 *int, x2 *int) {
    *x1, *x2 = *x2, *x1
    fmt.Println(*x1, *x2)
}

func main() {
	var a, b int 
	fmt.Scan(&a, &b)
	test(&a, &b)
}