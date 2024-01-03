//You must call the divide function which divides two numbers, but it returns not only the result of the division, but also error information.
//In case of any error you should output "error", if there is no error - the result of the function.
//The function divide(a int, b int) (int, error) receives two numbers as input that need to be divided and returns the result (int) and an error (error).

package main

import (
	"errors"
	"fmt"
)

func divide(a int, b int) (int, error) {
	if b == 0 {
		return -1, errors.New("division by zero")
	} 
	return a / b, nil
}

func main() {
    var a, b int 
    fmt.Scan(&a, &b)
    result, err := divide(a, b)
    if err != nil {
        fmt.Println(result, err)
    } else {
        fmt.Println(result)
    }
}