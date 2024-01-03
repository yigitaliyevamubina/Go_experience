//cityPopulation needs to be corrected so that it only stores information about cities from the groupCity[100] group.

package main

import "fmt"

func main() {
	groupcity := map[int][]string{
		10:   {"A", "B", "C"},
		100:  {"D", "E", "F"},
		1000: {"G", "H", "I"},
	}

	cityPopulation := map[string]int{
		"D": 100,
		"G": 1000,
		"A": 10,
		"F": 100,
	}

	for _, city := range append(groupcity[10], groupcity[1000]...) {
		delete(cityPopulation, city)
	}

	fmt.Println(cityPopulation)
}
