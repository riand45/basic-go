package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	firstName := "John"
	sayHelloTo(firstName, "Doe")
	sayHelloTo("Lorem", "Ipsum")
}