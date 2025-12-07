package main

import (
	"gopherstruct/list"
)

func main() {
	myList := list.LinkedList{}

	nodeOne := &list.Node{Data: 1}

	myList.Prepend(nodeOne)
	myList.Prepend(&list.Node{Data: 48})
	myList.Prepend(&list.Node{Data: 18})
	myList.Prepend(&list.Node{Data: 16})
	myList.Prepend(&list.Node{Data: 2})

	myList.PrintData()

	// test cases
	myList.DeleteByValue(16)
	myList.DeleteByValue(100)
	myList.DeleteByValue(2)

	emptyList := list.LinkedList{}
	emptyList.DeleteByValue(10)

	myList.PrintData()
}
