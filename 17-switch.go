package main

import "fmt"

func main() {

	name := "Johnasdadasd"

	switch name {
	case "John":
		fmt.Println("Hello John")
		fmt.Println("Hello John")
	case "Doe":
		fmt.Println("Hello Doe")
		fmt.Println("Hello Doe")
	default:
		fmt.Println("How are you?")
		fmt.Println("How are you")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("The name more 10 characters.")
	case length > 5:
		fmt.Println("The name more 5 characters")
	default:
		fmt.Println("The name passed")
	}
}