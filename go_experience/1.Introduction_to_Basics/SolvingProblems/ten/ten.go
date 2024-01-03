//По данному числу n закончите фразу "На лугу пасется..." одним из возможных продолжений: 
//"n коров", "n корова", "n коровы", правильно склоняя слово "корова".

package main 

import "fmt"

func main() {
	var n int 
	fmt.Scan(&n)
	if n == 1 {
		fmt.Println(n, "korova")
	} else if n >= 5 && n <= 20 {
		fmt.Println(n, "korov")
	} else if n % 10 == 1 {
		fmt.Println(n, "korova")	
	} else if n % 10 >= 2 && n % 10 <= 4{
		fmt.Println(n, "korovy")
	} else {
		fmt.Println(n, "korov")
	}
}