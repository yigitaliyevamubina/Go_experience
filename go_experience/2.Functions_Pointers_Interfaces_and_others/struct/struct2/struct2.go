package main

import (
  "fmt"
)

type Student struct {
  FullName         string 
  Age              int 
  MathScore        int
  EnglishScore     int 
  ProgrammingScore int 
  IsOnline         bool
}

func (s *Student) ScoreOverall() int {
  overall := (s.MathScore + s.EnglishScore + s.ProgrammingScore) / 3
  return overall
}

func main() {
  student := Student{
    FullName:         "Name Surname",
    Age:         15,
    EnglishScore:     100, 
    ProgrammingScore: 100,
    MathScore:       100,
    IsOnline:         false,

  }

  score := student.ScoreOverall()

  fmt.Println(score)
  fmt.Println(student)
}