package linkedlist

import "testing"

func TestLinkListCyle(t *testing.T) {

}

// hash算法实现链表回环检查
func hasCycle(head *ListNode) bool {
	m := map[*ListNode]struct{}{}

	for head != nil {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = struct{}{}
		head = head.Next
	}

	return false
}

type ListNode struct {
	Val  int
	Next *ListNode
}
