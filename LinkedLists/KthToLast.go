package main

import (
	"fmt"
)

/*
Return Kth to Last: Implement an algorithm to find the kth to last element of a singly linked list
*/

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
