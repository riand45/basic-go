package main

import (
"fmt"
"strings"
)

func main() {
	fmt.Println(strings.Contains("John Doe", "John"))
	fmt.Println(strings.Contains("John Doe", "Ashley"))

	fmt.Println(strings.Split("John Doe Lorem", " "))

	fmt.Println(strings.ToLower("John Doe Lorem"))
	fmt.Println(strings.ToUpper("John Doe Lorem"))
	fmt.Println(strings.ToTitle("John Doe Lorem"))

	fmt.Println(strings.Trim("      John Doe     ", " "))
	fmt.Println(strings.ReplaceAll("John Doe", "John", "Ashley"))
}