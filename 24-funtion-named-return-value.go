package main

import "fmt"

func getFullName2() (string, string, string) {
	return "Lorem", "Ipsum", "Dolor"
}

func main() {
	firstName, middleName, lastName := getFullName2()
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}