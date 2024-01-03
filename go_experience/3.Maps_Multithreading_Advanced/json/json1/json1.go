package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type NameInfo struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Length   int    `json:"length"`
	Language string `json:"lang"`
}

type Person struct {
	FullName    NameInfo
	Age         int
	Proffession string `json:"Current_job"`
	Birth_year  int
	Salary      float64
}

type Company struct {
	Name          string `json:"Company_name"`
	Launched_year int
	Fund          float64
	Director_name string   `json:"Boss"`
	Workers       []Person `json:"employers"`
}

type Country struct {
	Country_name string `json:"name"`
	Companies    []Company
	Capital      string
}

func main() {
	country1 := Country{
		Country_name: "Uzbekistan",
		Companies: []Company{
			{
				Name:          "Uzum",
				Launched_year: 2005,
				Fund:          123123.134,
				Director_name: "Akmal Rahimov",
				Workers: []Person{
					{
						FullName: NameInfo{
							Name:     "Ali",
							Surname:  "Valiyev",
							Length:   3,
							Language: "uz",
						},
						Age:         30,
						Proffession: "Enginer",
						Birth_year:  2000,
						Salary:      3425234524.34,
					},
					{
						FullName: NameInfo{
							Name:     "Vali",
							Surname:  "Aliyev",
							Length:   4,
							Language: "uz",
						},
						Age:         35,
						Proffession: "Manager",
						Birth_year:  1986,
						Salary:      70000.0,
					},
					{
						FullName: NameInfo{
							Name:     "Rashid",
							Surname:  "Olimov",
							Length:   6,
							Language: "uz",
						},
						Age:         28,
						Proffession: "Developer",
						Birth_year:  1993,
						Salary:      900.0,
					},
					{
						FullName: NameInfo{
							Name:     "Olim",
							Surname:  "Jasurov",
							Length:   4,
							Language: "uz",
						},
						Age:         32,
						Proffession: "Programmer",
						Birth_year:  2005,
						Salary:      90090.0,
					},
				},
			},
			{
				Name:          "Payme",
				Launched_year: 2000,
				Fund:          1663.00,
				Director_name: "Akmal Jasurov",
				Workers: []Person{
					{
						FullName: NameInfo{
							Name:     "Rustam",
							Surname:  "Alijonov",
							Length:   6,
							Language: "uz",
						},
						Age:         30,
						Proffession: "Waiter",
						Birth_year:  2002,
						Salary:      999.99,
					},
					{
						FullName: NameInfo{
							Name:     "John",
							Surname:  "Smith",
							Length:   4,
							Language: "en",
						},
						Age:         35,
						Proffession: "Businessman",
						Birth_year:  1997,
						Salary:      7122.0,
					},
					{
						FullName: NameInfo{
							Name:     "Davron",
							Surname:  "Olimov",
							Length:   6,
							Language: "uz",
						},
						Age:         28,
						Proffession: "Backend developer",
						Birth_year:  2004,
						Salary:      10000000.0,
					},
					{
						FullName: NameInfo{
							Name:     "Kamol",
							Surname:  "Anvarov",
							Length:   5,
							Language: "uz",
						},
						Age:         32,
						Proffession: "Manager",
						Birth_year:  1999,
						Salary:      1212321.123,
					},
				},
			},
		},
		Capital: "Tashkent",
	}

	js, err := json.MarshalIndent(&country1, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(js))

	file, err := os.Create("file.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	writer := bufio.NewWriter(file)
	writer.WriteString(string(js))

	writer.Flush()
	file.Close()

	country2 := Country{}

	file2, err := os.Open("file.json")
	if err != nil {
		fmt.Println(err)
		log.Fatal(0)
	}

	defer file2.Close()
	json2, err := os.ReadFile(file2.Name())
	if err != nil {
		fmt.Println(err)
		log.Fatal(0)
	}

	if !json.Valid(json2) {
		fmt.Println("Invalid json format")
		log.Fatal(0)
	}
	err3 := json.Unmarshal(json2, &country2)

	if err3 != nil {
		fmt.Println(err3)
		log.Fatal(0)
	}
	fmt.Println(country2)

}
