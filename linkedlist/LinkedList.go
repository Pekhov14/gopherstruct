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

	prev := l.head

	for prev.next != nil && prev.next.data != value {
		prev = prev.next
	}

	if prev.next == nil {
		return
	}

	// unlink current node from the list
	prev.next = prev.next.next
	l.lenth--
}
