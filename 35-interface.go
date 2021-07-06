package main

import "fmt"

type HasName interface {
	GetName() string
	GetAge() int
}

func yourName(hasName HasName) {
	fmt.Println("Hi", hasName.GetName(), ", your age", hasName.GetAge())
}

type Person struct {
	Name string
	Age int
}

type Animal struct {
	Name string
	Age int
}

func (person Person) GetName() string {
	return person.Name
}

func (person Person) GetAge() int {
	return person.Age
}

func (animal Animal) GetName() string {
	return animal.Name
}

func (animal Animal) GetAge() int {
	return animal.Age
}

func main() {
	var john Person
	john.Name = "John"
	john.Age = 30
	yourName(john)

	var cat Animal
	cat.Name = "Empus"
	cat.Age = 2
	yourName(cat)
}
