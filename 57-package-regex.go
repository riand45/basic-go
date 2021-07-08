package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("j([a-z])n")

	fmt.Println(regex.MatchString("john"))
	fmt.Println(regex.MatchString("honj"))
	fmt.Println(regex.MatchString("honj"))

	search := regex.FindAllString("john honj honj jh onhn johd", -1)
	fmt.Println(search)
}