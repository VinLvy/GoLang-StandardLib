package main

import (
	"container/list"
	"fmt"
)

func liss() {
	data := list.New()
	data.PushBack("daoa")
	data.PushBack("dapa")
	data.PushBack("dafa")

	var head *list.Element = data.Front()
	fmt.Println(head.Value) //daoa

	next := head.Next()
	fmt.Println(next.Value) //dapa

	next = next.Next()
	fmt.Println(next.Value) //dafa

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
