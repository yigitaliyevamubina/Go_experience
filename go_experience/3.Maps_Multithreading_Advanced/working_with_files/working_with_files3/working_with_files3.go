//This task will help you understand the encoding/csv package and path/filepath, although the archive/zip package can also be used for the solution (since the task file is provided in this format).

// The test archive, which you can download from our repository on github.com, contains a set of folders and files. One of these files is a data file in CSV format, while the other files do not contain structured data.

// You need to find and read this single file with structured data (this is a 10x10 table, separated by a comma), and as an answer you need to indicate the number located on the 5th line and 3rd position (indexes 4 and 2, respectively).

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func walkFunc(path string, info os.FileInfo, err error) error {

	if err != nil {
		return err
	}

	if info.IsDir() {
		return nil
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	content, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return err
	}

	if len(content) == 10 && len(content[4]) >= 3 {
		fmt.Println(path, info.Name(), content[4][2])
		fmt.Println(content)
	}

	return nil
}

func main() {
	const root = "./task"
	if err := filepath.Walk(root, walkFunc); err != nil {
		fmt.Println(err)
	}
}
