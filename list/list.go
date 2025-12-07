package list

import "fmt"

// TODO: write test for ll

// Node TODO: Make with generics, not only int.
type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l LinkedList) PrintData() {
	toPrint := l.head

	for l.length != 0 {
		fmt.Printf("%d ", toPrint.Data)
		toPrint = toPrint.Next
		l.length--
	}

	fmt.Println("\n")
}

// maybe rename to push
func (l *LinkedList) Prepend(n *Node) {
	second := l.head
	l.head = n
	l.head.Next = second
	l.length++
}

func (l *LinkedList) DeleteByValue(value int) {
	if l.length == 0 {
		return
	}

	// delete head
	if l.head.Data == value {
		l.head = l.head.Next
		l.length--
		return
	}

	prevToDelete := l.head

	for prevToDelete.Next.Data != value {
		if prevToDelete.Next.Next == nil {
			// value not found in ll
			return
		}

		prevToDelete = prevToDelete.Next
	}
	prevToDelete.Next = prevToDelete.Next.Next
	l.length--

}
