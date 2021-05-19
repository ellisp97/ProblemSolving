package main

import "fmt"

func main() {

	node_delete := TestDeletion()
	Display(node_delete)

	k := 3
	kth_node := _GetKthToLastNode(node_delete, k)
	fmt.Printf("%d'th node is %d \n", k, kth_node.data)

	can_delete := DeleteMiddleNode(kth_node.next.next)
	fmt.Printf("middle node delete returned %v \n", can_delete)

	TestPartitions()
}

func Display(n *Node) {
	for n != nil {
		fmt.Printf("%d \n", n.data)
		n = n.next
	}
}
