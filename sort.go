package main

import (
	"fmt"
	"sort"
)

type User struct {
	name string
	age  int
}

type UserSlice []User

func (userslice UserSlice) Len() int {
	return len(userslice)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].age < s[j].age
}

func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]

	/*
		temp := s[i]
		s[i] = s[j]
		s[j] = temp
	*/
}

func sorr() {
	users := []User{
		{"daoa", 16},
		{"budi", 100},
		{"doni", 50},
		{"dapa", 40},
		{"dono", 30},
	}
	sort.Sort(UserSlice(users))
	fmt.Println(users)
}
