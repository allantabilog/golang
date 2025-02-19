package main

import (
	"container/list"
)

func main() {
	list := list.New()

	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)

	for e := list.Front(); e != nil; e = e.Next() {
		println(e.Value.(int))
	}
}