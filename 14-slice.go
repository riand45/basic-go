package main

import "fmt"

func main()  {
	var month = [...]string {
		"January",
		"February",
		"Market",
		"April",
		"Mei",
		"June",
		"Juli",
		"Augustus",
		"September",
		"October",
		"November",
		"December",
	}
	//pointer = 4, length = 3, capacity = 8
	//slice := month[4:7]
	//fmt.Println(slice)

	//months[5] = "Change"
	//fmt.Println(slice1)

	//slice1[0] = "May"
	//fmt.Println(months)

	var slice2 = month[11:]
	//fmt.Println(slice2)

	var slice3 = append(slice2, "Month")
	fmt.Println(slice3)
	slice3[1] = "Not December"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(month)

	newSlice := make([]string, 2, 5)

	newSlice[0] = "John"
	newSlice[1] = "Doe"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}