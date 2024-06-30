package main

import (
	"fmt"
	"strconv"
)

func strc() {
	result, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("error", err.Error())
	} else {
		fmt.Println(result)
	}

	resultInt, err := strconv.Atoi("1000")
	if err != nil {
		fmt.Println("error", err.Error())
	} else {
		fmt.Println(resultInt)
	}

	binary := strconv.FormatInt(999, 2)
	fmt.Println(binary)

	stringInt := strconv.Itoa(999)
	fmt.Println(stringInt)
}
