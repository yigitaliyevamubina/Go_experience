package main

import (
  "fmt"
  "time"
)

func main() {
  fmt.Println(1*time.Hour + 3*time.Minute + 4*time.Second)
  fmt.Println(50*time.Millisecond)
  fmt.Println(time.April.String())

  t1 := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
  t2 := time.Date(2023, 1, 1, 17, 0, 0, 0, time.UTC)
  difference := t2.Sub(t1)
  fmt.Printf("Difference: %v\n", difference)
}