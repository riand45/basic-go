package main

import "fmt"

func main() {
	var password = "ZPassword"
	var repassword = "RPassword"
	var result bool = password == repassword
	fmt.Println(result)

	var value1 = 100
	var value2 = 200
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
