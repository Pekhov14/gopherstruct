package main

import "fmt"

// TODO: write test for ll

// Node TODO: Make with generics, not only int.
type Node struct {
	data int
	next *Node
}

type linkedList struct {
	head   *Node
	length int
}

func (l linkedList) printData() {
	toPrint := l.head

	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}

	fmt.Println("\n")
}

func (l *linkedList) prepend(n *Node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l *linkedList) deleteByValue(value int) {
	if l.length == 0 {
		return
	}

	// delete head
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	prevToDelete := l.head

	for prevToDelete.next.data != value {
		if prevToDelete.next.next == nil {
			// value not found in ll
			return
		}

		prevToDelete = prevToDelete.next
	}
	prevToDelete.next = prevToDelete.next.next
	l.length--

}

func main() {
	myList := linkedList{}

	nodeOne := &Node{data: 1}

	myList.prepend(nodeOne)
	myList.prepend(&Node{data: 48})
	myList.prepend(&Node{data: 18})
	myList.prepend(&Node{data: 16})
	myList.prepend(&Node{data: 11})
	myList.prepend(&Node{data: 7})
	myList.prepend(&Node{data: 2})

	myList.printData()

	// test cases
	myList.deleteByValue(16)
	myList.deleteByValue(100)
	myList.deleteByValue(2)

	emptyList := linkedList{}
	emptyList.deleteByValue(10)

	myList.printData()
}
