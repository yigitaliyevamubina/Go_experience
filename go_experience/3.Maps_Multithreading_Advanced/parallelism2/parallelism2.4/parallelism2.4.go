//example with ticker (returns data every n time)

package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.Tick(time.Second)
	cnt := 0 
	for {
		<-ticker
		fmt.Printf("%d tick\n", cnt + 1)
		cnt++
		if cnt == 5 {
			break
		}
	}
}