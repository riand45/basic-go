package main

import "fmt"

func main() {
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Looping to", counter)
	}

	slice := []string{"Lorem", "Ipsum", "Dolor", "Sit", "Amet"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	person := make(map[string]string)
	person["name"] = "John"
	person["title"] = "Doe"
	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}