package rll

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prevLink *ListNode
	for {
		if head == nil {
			break
		}
		next := head.Next
		head.Next = prevLink
		head = next
	}
	return head
}

func pr(slice []string) {
	fmt.Printf("cap: %d; len %d; slice: %v\n", cap(slice), len(slice), slice)
}

func slicelc() bool {
	var slice = make([]string, 0, 3)
	slice = append(slice, "hello", "world")
	// slice = []string{"hello", "world"}
	pr(slice)
	mod(slice)
	pr(slice)
	add(slice)
	pr(slice)
	return true
}

func mod(slice []string) {
	slice[0] = "googdbie"
}

func add(slice []string) {
	slice = append(slice, "!")
}
