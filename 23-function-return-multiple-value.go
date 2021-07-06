package main

import "fmt"

func getFullName() (firstName string, middleName string, lastName string) {
	firstName = "Lorem"
	middleName = "Ipsum"
	lastName = "Dolor"

	return
}

func main() {
	a, b, c := getFullName()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}