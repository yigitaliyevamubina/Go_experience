//you get 3 values ​​of type empty interface: if everything is successful, then the first 2 values ​​will be float64. 
//The third value will ideally be a string that can have the values: "+", "-", "*", "/" (a specific mathematical operation). 
//But such ideal cases will not always exist; you may receive values ​​of other types, and the string in the third value may not belong to one of the specified mathematical operations.

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func readTask() (interface{}, interface{}, interface{}) {

	var value1, value2, operation interface{}

	decoder := json.NewDecoder(os.Stdin)

	decoder.Decode(&value1)

	decoder.Decode(&value2)

	decoder.Decode(&operation)

	return value1, value2, operation

}

func main() {
	a, b, c := readTask()

	if _, ok := a.(float64); !ok {
		fmt.Printf("value=%v: %T", a, a)
		return
	}

	if _, ok := b.(float64); !ok {
		fmt.Printf("value=%v: %T", b, b)
		return
	}

	if val, ok := c.(string); ok {
		switch val {
		case "+":
			fmt.Printf("%.4f", a.(float64)+b.(float64))
		case "-":
			fmt.Printf("%.4f", a.(float64)-b.(float64))
		case "*":
			fmt.Printf("%.4f", a.(float64)*b.(float64))
		case "/":
			fmt.Printf("%.4f", a.(float64)/b.(float64))
		default:
			fmt.Println("неизвестная операция")
		}
	} else {
		fmt.Println("неизвестная операция")
	}
}
