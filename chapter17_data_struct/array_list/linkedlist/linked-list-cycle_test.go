package linkedlist

import "testing"

func TestLinkListCyle(t *testing.T) {

}

// 使用快慢表实现链表回环检查
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

// hash算法实现链表回环检查
func hasCycle1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

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
