package main

import (
	"fmt"
	"strings"
)

func stri() {
	fmt.Println(strings.Contains("daoa dapa", "daoa"))
	fmt.Println(strings.Split("daoa dapa", " "))
	fmt.Println(strings.ToLower("Daoa Dapa"))
	fmt.Println(strings.ToUpper("daoa dapa"))
	fmt.Println(strings.Trim("     daoa dapa    ", " "))
	fmt.Println(strings.ReplaceAll("daoa rendra", "daoa", "dapa"))
}
