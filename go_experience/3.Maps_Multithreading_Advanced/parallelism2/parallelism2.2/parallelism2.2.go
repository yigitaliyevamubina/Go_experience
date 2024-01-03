//You need to call the work() function in a separate goroutine and wait for the results of its execution.

// The work() function does not accept or return anything.

package main

import (
	"fmt"
	"time"
)

func work() {
	time.Sleep(3 * time.Millisecond)
}

func main() {
	go work()
	time.Sleep(3 * time.Millisecond)
	fmt.Println("Done")
}
