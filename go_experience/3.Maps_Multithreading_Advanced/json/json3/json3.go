//This task is focused on real work with data in JSON format. To solve this, we will use the OKVED (All-Russian Classifier of Types of Economic Activities) directory, which was posted on the web portal data.gov.ru.

// The data structure information you need is contained in the file structure-20190514T0000.json, and the data you need to decode is contained in the file data-20190514T0100.json. The files are hosted in our repository on github.com.

// In order to show that you were really able to decode the document, you need to write down as a response the sum of the global_id fields of all elements encoded in the file.

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Datastruct struct {
	GloablId int `json:"global_id"`
}

type Datamassiv struct {
	elements []Datastruct
}

func main() {
	jsonFile := "text.json"
	js, _ := os.ReadFile(jsonFile)

	var data Datamassiv
	err := json.Unmarshal(js, &data.elements)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)
	sum := 0
	for i := range data.elements {
		sum += data.elements[i].GloablId
	}
	fmt.Println(sum)
}
