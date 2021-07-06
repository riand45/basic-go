package main

import (
	"errors"
	"fmt"
)

func Divider(value1 int, value2 int) (int, error) {
	if value1 == 0 || value2 == 0 {
		return 0, errors.New("Value can't zero")
	} else {
		result := value1 / value2
		return result, nil
	}
}

func main() {
	result, err := Divider(10, 10)
	if err == nil {
		fmt.Println("Result", result)
	} else {
		fmt.Println("Error", err.Error())
	}
}
