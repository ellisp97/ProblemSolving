package main

/*
Partition a linked list around a value x, such that all nodes less than x come before all nodes greater than or equal to x
partition element x can appear anywhere in right partition

input: 3 -> 5 -> 8 -> 5 -> 10 -> 2 -> 1 [partition: 5]
output:3 -> 1 -> 2      ->       10 -> 5 -> 5 -> 8

*/

func PartitionLinkedList(node *Node, partition_val int) *Node {
	head := node
	tail := node
	for n := node.next; n != nil; n = n.next {
		temp := &Node{data: n.data}

		if temp.data < partition_val {
			temp.next = head
			head = temp
		} else {
			tail.next = temp
			tail = tail.next
		}
	}
	return head

}

func TestPartitions() {
	n := &Node{data: 3}
	appendToTail(n, 5)
	appendToTail(n, 8)
	appendToTail(n, 5)
	appendToTail(n, 10)
	appendToTail(n, 2)
	appendToTail(n, 1)

	n1 := PartitionLinkedList(n, 5)
	Display(n1)
}
