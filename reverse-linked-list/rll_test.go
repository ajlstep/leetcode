package rll

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPolindrome(t *testing.T) {
	s := createNodes([]int{1, 3, 6, 8})
	r := createNodes([]int{8, 6, 3, 1})
	require.Equal(t, reverseList(s), r, fmt.Sprintf("%v|%v", s, r))
}

func createNodes(arr []int) *ListNode {
	prevListNode := &ListNode{}
	for i := len(arr) - 1; i >= 0; i-- {
		nowListNode := ListNode{
			Val:  arr[i],
			Next: prevListNode,
		}
		prevListNode = &nowListNode
	}
	return prevListNode
}

func TestSlicelc(t *testing.T) {
	a := slicelc()
	require.Equal(t, slicelc(), a, "test")
}
