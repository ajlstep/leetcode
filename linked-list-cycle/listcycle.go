package linkedlistcycle

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

func hasCycle(head *ListNode) bool {
	fmt.Printf("%p", head)
	var mp = make(map[string]interface{})
	now := head
	for {
		if now.Next == nil {
			break
		}
		if _, found := mp[fmt.Sprintf("%p", now)]; found {
			return true
		}
		mp[fmt.Sprintf("%p", now)] = nil
		now = now.Next
	}
	return false
}
