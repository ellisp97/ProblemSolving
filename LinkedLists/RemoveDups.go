package main

func appendToTail(n *Node, d int) {
	end := &Node{data: d}
	for n.next != nil {
		n = n.next
	}
	n.next = end
}

func RemoveDuplicatesFromLinkedList(n *Node) {
	if n == nil {
		return
	}
	set := make(map[int]bool)
	var previous *Node
	for n != nil {
		if set[n.data] {
			previous.next = n.next
		} else {
			set[n.data] = true
			previous = n
		}
		n = n.next
	}
}

func TestDeletion() *Node {

	node1 := Node{data: 1}

	int_arr := []int{2, 3, 4, 4, 5}

	for _, s := range int_arr {
		appendToTail(&node1, s)
	}
	RemoveDuplicatesFromLinkedList(&node1)
	return &node1
}
