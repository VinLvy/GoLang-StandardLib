package main

import (
	"fmt"
	"regexp"
)

func rege() {
	var regex *regexp.Regexp = regexp.MustCompile(`da([a-z])a`)

	fmt.Println(regex.MatchString("dapa"))
	fmt.Println(regex.MatchString("daoa"))
	fmt.Println(regex.MatchString("dafa"))

	fmt.Println(regex.FindAllString("dapa dafa daoa da1a daoo", 10))
}
