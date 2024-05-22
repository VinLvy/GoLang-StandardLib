package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func read() {
	csvstring := "daoa, dapa, dafa\n" +
		"budi, dono, doni\n" +
		"indo, jepang, peru"

	reader := csv.NewReader(strings.NewReader(csvstring))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}
}
