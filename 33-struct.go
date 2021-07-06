
package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var john Customer
	john.Name = "John"
	john.Address = "Indonesia"
	john.Age = 30

	//fmt.Println(eko.Name)
	//fmt.Println(eko.Address)
	//fmt.Println(eko.Age)
	//
	doe := Customer{
		Name:    "Doe",
		Address: "United State",
		Age:     35,
	}
	fmt.Println(doe)
	//
	//doe := Customer{"Doe", "Jakarta", 22}
	//fmt.Println(doe)
}