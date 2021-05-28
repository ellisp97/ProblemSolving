package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

/*
Implement a function to check if a linkedlist is a palindrome

palindrome is the same backwards e.g. racecar
conditions of a palindrome has to be either:
	- even number of letters with identical counts
	- odd number with only one letter has count 1
*/

//  ========== Method1: Reverse and Compare ================
func CheckPalindrome(node *Node) bool {
	reversed_node := Reverse(node)
	return isEqual(reversed_node, node)
}

func Reverse(node *Node) *Node {
	var head *Node
	for node != nil {
		n := &Node{data: node.data}
		n.next = head
		head = n
		node = node.next
	}
	return head
}

func isEqual(n1 *Node, n2 *Node) bool {
	for n1 != nil && n2 != nil {
		if n1.data != n2.data {
			return false
		}
		n1 = n1.next
		n2 = n2.next
	}

	// If one is longer than the other they cant be equal
	if n1 == nil && n2 == nil {
		return true
	} else {
		return false
	}
}

// ========================================================

// ============ Method2: Iterative - Multi Runner Approach =================

// Can use stack to check if front half is reverse of second half

func CheckPalindrome2(node Node) bool {
	fast_runner := &node
	slow_runner := &node

	first_half_stack := stack.New()

	// When fast runner reaches end (as we're moving 2x as quick) slow runner is at middle
	for fast_runner != nil && fast_runner.next != nil {
		first_half_stack.Push(slow_runner.data)
		slow_runner = slow_runner.next
		fast_runner = fast_runner.next.next
	}

	// If odd there will be one more node left
	if fast_runner != nil {
		slow_runner = slow_runner.next
	}

	// Travel through second half of the list making a comparison to first half stored in the stack
	for slow_runner != nil {
		if first_half_stack.Pop() != slow_runner.data {
			return false
		}
		slow_runner = slow_runner.next
	}
	return true
}

// =========================================================================

func TestPalindrome() {
	n1 := &Node{data: 2}
	appendToTail(n1, 4)
	appendToTail(n1, 1)
	appendToTail(n1, 4)
	appendToTail(n1, 2)

	fmt.Println(CheckPalindrome(n1))
	fmt.Println(CheckPalindrome2(*n1))
}
