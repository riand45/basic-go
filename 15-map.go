package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "John Doe",
		"address": "United State",
	}
	person["name"] = "Ahmed"
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Lear Go-Lang"
	book["author"] = "Google"
	book["ups"] = "Salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))
}