package main

import (
	"errors"
	"fmt"
)

var (
	Validationerror = errors.New("Valditaion error")
	Notfounderror   = errors.New("Not found error")
)

func GetById(id string) error {
	if id == "" {
		return Validationerror
	}
	if id != "daoa" {
		return Notfounderror
	}
	// sukses
	return nil
}

func erro() {
	err := GetById("dao")
	if err != nil {
		if errors.Is(err, Validationerror) {
			fmt.Println("Validation error")
		} else if errors.Is(err, Notfounderror) {
			fmt.Println("Not found error")
		} else {
			fmt.Println("Unknown error")
		}
	}
}
