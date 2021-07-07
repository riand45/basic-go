package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age int
}

type UserSlice []User

func (item UserSlice) Len() int {
	return len(item)
}

func (item UserSlice) Less(i, j int) bool {
	return item[i].Age < item[j].Age
}

func (item UserSlice) Swap(i, j int) {
	item[i], item[j] = item[j], item[i]
}

func main() {
	users := []User {
		{"Lorem", 30},
		{"Ipsum", 35},
		{"Dolor", 31},
		{"Sit Amet", 25},
	}

	fmt.Println(users)
	sort.Sort(UserSlice(users))
	fmt.Println(users)
}