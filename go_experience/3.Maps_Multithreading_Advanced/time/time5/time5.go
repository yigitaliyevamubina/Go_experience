package main

import (
  "fmt"
  "time"
)

func main() {
  loc, _ := time.LoadLocation("Asia/Uzbekistan")
  fmt.Println(loc)
  fmt.Println(time.Now())

  currentTime := time.Date(2023, time.December, 22, 12, 27, 30, 55, time.Local)

  fmt.Println(currentTime.Date())
  fmt.Println(currentTime.AddDate(2, 2, 2))
  fmt.Println(currentTime.Date())
  fmt.Println(currentTime.Year())
  fmt.Println(currentTime.Month())
  fmt.Println(currentTime.Day())
  fmt.Println(currentTime.Hour())
  fmt.Println(currentTime.Minute())
  fmt.Println(currentTime.Second())
  fmt.Println(currentTime.Unix())
  fmt.Println(currentTime.Weekday())
  fmt.Println(currentTime.YearDay())

  firstTime := time.Date(2022, time.January, 23, 23, 11, 34, 10, time.UTC)
  secondTime := time.Date(2024 ,time.February, 10, 16, 45, 15, 19, time.Local)

  fmt.Println(currentTime.Format("02-01-2006 15:04:05"))

  fmt.Println(firstTime.Equal(secondTime))
  fmt.Println(firstTime.After(secondTime))
  fmt.Println(secondTime.Before(firstTime))

  thirdTime := time.Date(1999, time.July, 19, 18, 25, 15, 59, time.Local)

  changedTime := thirdTime.AddDate(-2, -1, -1)
  fmt.Println(thirdTime.Sub(changedTime))

}

