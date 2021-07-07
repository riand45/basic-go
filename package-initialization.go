package main

import (
	"belajar/database"
	"fmt"
)

func main() {
	result := database.GetDatabase()

	fmt.Println(result)
}
