package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		//Odd
		//if i % 2 == 0 {
		//	continue
		//}
		//Even
		if i % 2 == 1 {
			continue
		}

		if i >= 5 {
			break
		}

		fmt.Println("Looping to", i)
	}
}
