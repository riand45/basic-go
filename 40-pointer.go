package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Sidoarjo", "Jawa Timur", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1

	address2.Province = "Jawa Kentel"

	*address2 = Address{"Bandung", "Jawa Barat", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address)
	address4.City = "Bekasi"
	fmt.Println(address4)
}
