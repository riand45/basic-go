package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string {
			"name": name,
		}
	}
}

func main() {
	var person = newMap("aa")

	if person == nil {
		fmt.Println("Data empty")
	} else {
		fmt.Println(person)
	}
}