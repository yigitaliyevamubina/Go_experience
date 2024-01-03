package main

import (
	"fmt"
	"strconv"
)

func main() {

	var bigint int64 = 130
	var little = int8(bigint)
	fmt.Println(little)

	var num float64 = 12.99
	var intNum int64 = int64(num)
	fmt.Println(intNum)

	a := "hello"
	b := []rune(a)
	c := string(b)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(strconv.FormatInt(int64(100), 8))

	number := 12.2
	fmt.Println(strconv.FormatFloat(number, 'f', -1, 64))

	d := "15"
	e := "10"
	f, err := strconv.Atoi(d)
	if err != nil {
		panic(err)
	}
	g, err := strconv.Atoi(e)
	if err != nil {
		panic(err)
	}
	fmt.Println(f - g)

	float_num := "15.1234"
	res, err := strconv.ParseFloat(float_num, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	boolian := "true"
	resp, err := strconv.ParseBool(boolian)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

}
