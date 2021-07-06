package main

import "fmt"

type Customers struct {
	Name, Address string
	Age           int
}

func (customer Customers) sayHi(name string){
	fmt.Println("Hi, ", name, "My Name is", customer.Name)
}

func (customer Customers) sayHello(){
	fmt.Println("Say Hello from", customer.Name)
}

func main() {
	var oneCustomer Customers

	oneCustomer.sayHi("John Doe")
	oneCustomer.sayHello()
}
