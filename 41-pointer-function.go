package main

import "fmt"

type Location struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(location *Location) {
	location.Country = "Indonesia"
}

func main() {
	var location = Location{
		City:     "Surabaya",
		Province: "Jawa Timur",
		Country:  "",
	}
	ChangeCountryToIndonesia(&location)
	fmt.Println(location)
}