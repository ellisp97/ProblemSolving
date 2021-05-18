package main

import "fmt"

func main() {

	node_delete := TestDeletion()
	Display(node_delete)

	k := 3
	kth_node := _GetKthToLastNode(node_delete, k)
	fmt.Printf("%d'th node is %d", k, kth_node.data)

}

func Display(n *Node) {
	for n != nil {
		fmt.Printf("%d \n", n.data)
		n = n.next
	}
}
