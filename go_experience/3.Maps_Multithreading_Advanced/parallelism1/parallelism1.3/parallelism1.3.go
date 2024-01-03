//Write a pipeline element (function) that remembers the previous value and sends the value to the next stage of the pipeline only if it is different from what came previously.

// Your function should accept two channels - inputStream and outputStream, in the first one you will receive strings, in the second you should send values ​​without repetitions.
//As a result, outputStream should contain values ​​that are not repeated in a row. Don't forget to close the channel ;)

// The function should be called removeDuplicates()

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func removeDuplicates(inputStream chan string, outputStream chan string) {
    s := ""
    for v := range inputStream {
		if strings.Contains(s, v) {
			continue
		} else {
			s += v 
			outputStream <- v
		}

    }
    close(outputStream)
}

func main() {
	inputStream := make(chan string, 10)
	outputStream := make(chan string, 10)

	go func() {
		for i := 0; i < 10; i++ {
			inputStream <- strconv.Itoa(rand.Intn(10))
		}
		close(inputStream)
	}()

	go removeDuplicates(inputStream, outputStream)

	for {
		if n, ok := <- outputStream; ok {
			fmt.Println(n)
		} else {
			fmt.Println("The end")
			break
		}
	}
}