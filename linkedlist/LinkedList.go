package linkedlist

type Node struct {
    data int
    next  *Node
}

type LinkedList struct {
	head  *Node
	lenth int
}

func (l *LinkedList) PushFront (n *Node) {
	tmp := l.head
	l.head = n
	l.head.next = tmp
	l.lenth++
}

func (l *LinkedList) removeByValue(value int) {
	if l.lenth == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.lenth--
		return
	}

	prevToRemove := l.head

	// prevToRemove.next is current node
	for prevToRemove.next.data != value {
		if prevToRemove.next.next == nil {
			return
		}

		prevToRemove = prevToRemove.next
	}

	prevToRemove.next = prevToRemove.next.next
	l.lenth--
}
