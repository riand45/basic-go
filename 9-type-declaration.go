package main

import "fmt"

func main() {
	type Name string
	type Age int
	type Status bool

	var name Name = "John Doe"
	var age Age = 31
	var status Status = true

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(status)
}
