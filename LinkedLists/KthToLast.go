package main

import (
	"fmt"
)

/*
Return Kth to Last: Implement an algorithm to find the kth to last element of a singly linked list
*/

type Node struct {
	data int
	next *Node
	prev *Node
	key  interface{}
}

func appendToTail(n *Node, d int) {
	end := &Node{data: d}
	for n.next != nil {
		n = n.next
	}
	n.next = end
}

// Option A: Just print the value

func PrintKthToLastNode(head *Node, k int) int {

	if head == nil {
		return 0
	}
	index := PrintKthToLastNode(head.next, k) + 1
	if index == k {
		fmt.Printf("%d 'th to last node is %d \n", k, head.data)
	}
	return index
}

// Option B: Use pointers to return whole node
func GetKthToLastNode(head *Node, k int, i *int) *Node {
	if head == nil {
		return nil
	}
	n := GetKthToLastNode(head.next, k, i)
	*i++
	if *i == k {
		return head
	}
	return n

}

func _GetKthToLastNode(head *Node, k int) *Node {
	i := 0
	return GetKthToLastNode(head, k, &i)
}

func main() {
	k := 2

	node1 := Node{data: 20}
	appendToTail(&node1, 2)
	appendToTail(&node1, 3)
	appendToTail(&node1, 4)

	PrintKthToLastNode(&node1, k)

	n := _GetKthToLastNode(&node1, k)
	fmt.Printf("%d 'th to last node is %d", k, n.data)
}
