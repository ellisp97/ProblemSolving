package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
	key  interface{}
}

func (n *Node) Len() int {
	var counter int
	for n != nil {
		counter++
		n = n.next
	}
	return counter
}

func appendToTail(n *Node, d int) {
	end := &Node{data: d}
	for n.next != nil {
		n = n.next
	}
	n.next = end
}

type LinkedList struct {
	length int
	head   *Node
	tail   *Node
}

func (l LinkedList) Len() int {
	return l.length
}

func (l LinkedList) Display() {
	for l.head != nil {
		fmt.Printf("%v -> ", l.head.data)
		l.head = l.head.next
	}
	fmt.Println()
}

func (l *LinkedList) PushBack(n *Node) {
	if l.head == nil {
		l.head = n
		l.tail = n
		l.length++
	} else {
		l.tail.next = n
		l.tail = n
		l.length++
	}
}

func (l *LinkedList) DeleteNode(key int) {
	if l.head.data == key {
		l.head = l.head.next
		l.length--
		return
	}
	var prev *Node = nil
	curr := l.head
	for curr != nil && curr.data != key {
		prev = curr
		curr = curr.next
	}

	if curr == nil {
		fmt.Println("Key not found")
		return
	}
	prev.next = curr.next
	l.length--
	fmt.Println("Node deleted")
}

// func main() {
// 	list := LinkedList{}
// 	node1 := &Node{data: 1}
// 	node2 := &Node{data: 2}
// 	node3 := &Node{data: 3}
// 	node4 := &Node{data: 4}
// 	node5 := &Node{data: 6}
// 	list.PushBack(node1)
// 	list.PushBack(node2)
// 	list.PushBack(node3)
// 	list.PushBack(node4)
// 	list.PushBack(node5)
// 	fmt.Println("Length = ", list.Len())
// 	list.Display()
// 	list.DeleteNode(4)
// 	fmt.Println("Length = ", list.Len())
// 	list.Display()
// }
