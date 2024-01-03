package main

import "fmt"

type test struct {
	On bool
	Ammo int
	Power int
}

func (t *test) Shoot() bool {
	if t.On && t.Ammo > 0{
		t.Ammo--
		return true
	} 
	return false
  }

func (t *test) RideBike() bool {
	if t.On && t.Power > 0{
		t.Power--
		return true
	} 
	return false
  }

func main() {
	testStruct := new(test)
	testStruct.Ammo = 10 
	testStruct.On = true
	testStruct.Power = 15
	fmt.Println(testStruct.RideBike())
	fmt.Println(testStruct.Shoot())
}
