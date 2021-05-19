package main

/*
Implement an algorithm to delete the middle node (any but first and last not neccess exact middle)
of a singly linked list given only access to that node

e.g. Input: node c from linked list a -> b -> c -> d -> e -> f
	Result: nothing returned, but the new linked list looks like a -> b -> d -> e -> f

*/

func DeleteMiddleNode(n *Node) bool {
	if n == nil || n.next == nil { // Either the head or the tail
		return false
	}

	next := n.next
	n.data = next.data
	n.next = next.next
	return true
}
