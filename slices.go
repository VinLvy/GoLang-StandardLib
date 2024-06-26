package main

import (
	"fmt"
	"slices"
)

func slic() {
	names := []string{"tony", "dono", "daoa", "dapa"}
	values := []int{100, 90, 80, 70, 60}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "daoa"))
	fmt.Println(slices.Index(names, "daoa"))
	fmt.Println(slices.Index(names, "tony"))
}
