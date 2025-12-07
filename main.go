package main

import (
	"gopherstruct/list"
)

func main() {
	myList := list.LinkedList{}

	nodeOne := &list.Node{Data: 1}

	myList.PushFront(nodeOne)
	myList.PushFront(&list.Node{Data: 48})
	myList.PushFront(&list.Node{Data: 18})
	myList.PushFront(&list.Node{Data: 16})
	myList.PushFront(&list.Node{Data: 2})

	myList.PrintData()

	// test cases
	myList.DeleteByValue(16)
	myList.DeleteByValue(100)
	myList.DeleteByValue(2)
	myList.PushBack(&list.Node{Data: 99})

	emptyList := list.LinkedList{}
	emptyList.DeleteByValue(10)
	emptyList.PrintData()

	myList.PrintData()
}
