package main

import "fmt"

func main() {
	name := "Lorem Ipsum"
	counter := 0

	increment := func() {
		name := "Dolor"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}