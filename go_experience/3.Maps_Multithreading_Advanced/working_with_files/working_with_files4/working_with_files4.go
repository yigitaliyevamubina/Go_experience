//This task is mainly focused on studying the bufio.Reader type, since this type allows you to read data gradually.

// The test file, which you can download from our github.com repository, contains a long series of numbers separated by a ";". You need to find what position the number 0 is in and indicate it as the answer. 
//It is required to display exactly the position of the number, and not the index (that is, the serial number, numbering starting from 1).

// For example: 12;234;6;0;78 :
// The correct answer will be the serial number of the number: 4

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("t.data")
	if err != nil {
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(file)
	indx := 1
	defer file.Close()

	for {
		line, err := reader.ReadString(';')
		if err != nil {
			break
		}

		fmt.Println(line)

		line = strings.Trim(line, ";")

		if strings.Contains(line, "0") && len(line) == 1 {
			fmt.Println(indx)
			return
		}
		indx++
	}
}
