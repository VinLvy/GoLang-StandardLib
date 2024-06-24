package main

import (
	"fmt"
	"path"
)

func patt() {
	fmt.Println(path.Dir("hello/world.go"))
	fmt.Println(path.Base("hello/world.go"))
	fmt.Println(path.Ext("hello/world.go"))
	fmt.Println(path.Join("hello", "world", "main.go"))
}
