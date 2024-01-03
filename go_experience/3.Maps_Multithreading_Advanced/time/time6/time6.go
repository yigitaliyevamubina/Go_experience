package main

import (
  "bufio"
  "fmt"
  "os"
  "time"
)

func main() {
  var t string
  reader := bufio.NewScanner(os.Stdin)
  reader.Scan()
  t = reader.Text()
  ans, _ := time.Parse(time.RFC3339, t)
  fmt.Println(ans.Format("Mon Jan 2 15:04:05 MST 2006"))

}
