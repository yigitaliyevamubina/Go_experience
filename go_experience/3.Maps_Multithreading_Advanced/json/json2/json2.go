//Data about university group students in JSON format is submitted to standard input.

//Information about each student contains information about the grades he received (Rating). You need to read the data and calculate the average number of marks received by the students in the group.

//The answer to the problem must be written to standard output in JSON format in the following form:
// {
//     "Average": 14.1
// }

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Student struct {
	Rating []int `json:"Rating"`
}

type AllStudents struct {
	Students []Student `json:"Students"`
}

type Rating struct {
	Average float32
}

func main() {
	js, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}

	var students AllStudents
	err2 := json.Unmarshal(js, &students)
	if err2 != nil {
		fmt.Println(err)
	}

	cnt1 := 0
	cnt2 := 0
	for i := range students.Students {
		cnt1++
		cnt2 += len(students.Students[i].Rating)
	}
	aver := Rating{
		Average: float32(cnt2) / float32(cnt1),
	}

	jsn, err3 := json.MarshalIndent(aver, "", "    ")
	if err3 != nil {
		fmt.Println(err3)
		return
	}
	os.Stdout.WriteString(string(jsn))
}
