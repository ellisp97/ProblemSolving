package main

import "fmt"

/*
Sum Lists: You have two nnumbers represented by a linked list, where each node contains a single digit
The digits are stored in _reverse_ order, such that the 1s digit is at the head of the list

Write a function that adds the two numbers and returns the sum as a linked list

e.g.

Input (7 -> 1 -> 6) + (5 -> 9 -> 2)  617 + 295
Output: (2 -> 1 -> 9) 912
*/

func addLists(n1 *Node, n2 *Node) *Node {
	return addListsWithCarry(n1, n2, 0)
}

func addListsWithCarry(n1 *Node, n2 *Node, carry int) *Node {
	if n1 == nil && n2 == nil && carry == 0 {
		return nil
	}

	var ret Node
	value := carry
	if n1 != nil {
		value += n1.data
	}

	if n2 != nil {
		value += n2.data
	}

	ret.data = value % 10 // Second digit added to current node, first left for carry
	if n1 != nil || n2 != nil {
		r_val := 0
		if value > 10 {
			r_val = 1
		}
		var remainder *Node
		if (n1 == nil || n1.next == nil) && (n2 == nil || n2.next == nil) {
			remainder = addListsWithCarry(nil, nil, r_val)
		} else if n1 == nil || n1.next == nil {
			remainder = addListsWithCarry(nil, n2.next, r_val)
		} else if n2 == nil || n2.next == nil {
			remainder = addListsWithCarry(n1.next, nil, r_val)
		} else {
			remainder = addListsWithCarry(n1.next, n2.next, r_val)
		}
		ret.next = remainder
	}
	return &ret
}

/*

Repeat the first problem but in forward order

e.g.

Input (6 -> 1 -> 7) + (2 -> 9 -> 5)  617 + 295
Output: (9 -> 1 -> 2) 912

Uneven length lists have to be padded
e.g.
	(1 -> 2 -> 3 -> 4) + (5 -> 6 -> 7)
	we need to know 5 matches to 2
	so pad list with 0
	0 -> 5 -> 6 -> 7
*/
type HelperClass struct {
	carry int
	n     *Node
}

func addListsBackwards(n1 *Node, n2 *Node) *Node {
	len1 := n1.Len()
	len2 := n2.Len()

	if len1 < len2 {
		n1 = padList(n1, len2-len1)
	} else if len2 > len1 {
		n2 = padList(n2, len1-len2)
	}

	sum := addListsRecursion(n1, n2)
	// sum will store value of final remainder in carry, if >0 add this
	if sum.carry == 0 {
		return sum.n
	} else {
		n := insertBefore(sum.n, sum.carry)
		return n
	}
}

func addListsRecursion(n1 *Node, n2 *Node) HelperClass {
	if n1 == nil && n2 == nil {
		sum := HelperClass{carry: 0, n: nil}
		return sum
	}

	// Add smaller digits
	sum := addListsRecursion(n1.next, n2.next)

	// Add carry to current data
	val := sum.carry + n1.data + n2.data

	// Insert sum of current digits
	res := insertBefore(sum.n, val%10)

	sum.n = res
	sum.carry = val / 10
	return sum
}

func padList(head *Node, padding int) *Node {
	for i := 0; i < padding; i++ {
		head = insertBefore(head, 0)
	}
	return head
}

func insertBefore(node *Node, insert int) *Node {
	n := &Node{data: insert}
	if node != nil {
		n.next = node
	}
	return n
}

func TestSums() {
	n1 := &Node{data: 7}
	appendToTail(n1, 1)
	appendToTail(n1, 6)
	appendToTail(n1, 3)

	n2 := &Node{data: 5}
	appendToTail(n2, 9)
	appendToTail(n2, 2)

	ans := addLists(n1, n2)
	fmt.Println("sum of the lists is")
	Display(ans)
}

func TestSumsBackwards() {
	n1 := &Node{data: 6}
	appendToTail(n1, 1)
	appendToTail(n1, 7)

	n2 := &Node{data: 2}
	appendToTail(n2, 9)
	appendToTail(n2, 5)

	ans := addListsBackwards(n1, n2)
	fmt.Println("sum of the lists is")
	Display(ans)
}
