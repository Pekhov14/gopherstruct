package linkedlist

import "fmt"

func (l LinkedList) dump() {
	for node := l.head; node != nil; node = node.next {
		fmt.Println(node.data)
	}
}

func RunTest() {
	customLinkedList := LinkedList{}

	node0 := &Node{data: 11}
	node1 := &Node{data: 22}
	node2 := &Node{data: 33}
	node3 := &Node{data: 44}
	node4 := &Node{data: 55}
	node5 := &Node{data: 66}
	node6 := &Node{data: 77}

	customLinkedList.PushFront(node0)
	customLinkedList.PushFront(node1)
	customLinkedList.PushFront(node2)
	customLinkedList.PushFront(node3)
	customLinkedList.PushFront(node4)
	customLinkedList.PushFront(node5)
	customLinkedList.PushFront(node6)

	customLinkedList.dump()
	customLinkedList.removeByValue(707)
	customLinkedList.removeByValue(77)
	customLinkedList.dump()

	emptyList := LinkedList{}
	emptyList.removeByValue(10)
}
