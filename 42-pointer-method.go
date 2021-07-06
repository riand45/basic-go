package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	person := Man{"John Doe"}
	person.Married()

	fmt.Println(person.Name)
}