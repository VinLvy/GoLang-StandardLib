package main

import "fmt"

func fmti() {
	fmt.Println("Hello world")

	firstname := "daoa"
	lastname := "dapa"

	fmt.Println("Hello '", firstname, lastname, "'")
	fmt.Printf("Hello '%s %s'\n", firstname, lastname)
}
