//The input is a float64 number. You need to output the converted number according to the rule: the number is squared, 
//then the fractional part is cut off so that 4 decimal places remain. But if the number is equal to or less than zero, output:
// "the number R is not suitable", where R is the original number with 2 digits after the decimal point and 2 in width.
// And if the number is more than 10,000, display the original number in exponential notation (the exponent sign is in lowercase).

// Sample input:
// -000012.2123
// Sample output:
// the number -12.21 is not suitable

package main

import (
  "fmt"
)

func main() {
  var n float64
  fmt.Scan(&n)

  if n <= 0 {
    fmt.Printf("%2.2f is not suitable", n)
  } else if n > 10000 {
    fmt.Printf("%e", n)
  } else {
    ans := n * n
    fmt.Printf("%.4f", ans)
  }
}