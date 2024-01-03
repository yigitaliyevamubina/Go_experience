package main

import (
	"fmt"
	"math"
)

var p float64= 1296
var v float64 = 6 
var k float64 = 6


func M() float64 {
	ans := p * v
	return ans
}

func W() float64 {
	m := M()
	ans := math.Sqrt(k / m)
	return ans
}

func T() float64 {
	w := W()
	return 6 / w
}

func main() {
	fmt.Println(T())
}