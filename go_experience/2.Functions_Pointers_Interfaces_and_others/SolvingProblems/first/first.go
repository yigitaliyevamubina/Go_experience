//The input is a and b - legs of a right triangle. We need to find the length of the hypotenuse.

package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	c := int(math.Sqrt(float64((a * a) + (b * b))))
	fmt.Println(c)
}