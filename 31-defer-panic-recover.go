package main

import "fmt"

func logging(){
	message := recover()
	if message != nil {
		fmt.Println(message)
	}
	fmt.Println("Finished calling function")
}

func runApplication(value int){
	defer logging()
	fmt.Println("Run application")

	if value == 0 {
		panic("Application Error")
	}

	result := 10 / value

	fmt.Println("Result", result)
}

func main() {
	runApplication(0)
}