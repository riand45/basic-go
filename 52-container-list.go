package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Lorem")
	data.PushBack("Ipsum")
	data.PushBack("Dolor")
	data.PushFront("Sit Amet")

	// From front to back
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	// From back to front
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}