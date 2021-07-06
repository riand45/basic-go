package main

import "fmt"

func Check(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return false
	}
}

func main() {
	var data interface{} = Check(2)

	fmt.Println(data)
}