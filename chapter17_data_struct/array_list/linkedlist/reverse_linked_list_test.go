package linkedlist

// 翻转链表
// 定义两个指针： prepre 和 curcur ；prepre 在前 curcur 在后。
// 每次让 prepre 的 nextnext 指向 curcur ，实现一次局部反转
// 局部反转完成之后，prepre 和 curcur 同时往前移动一个位置
// 循环上述过程，直至 prepre 到达链表尾部
//
//作者：huwt
//链接：https://leetcode.cn/problems/reverse-linked-list/solution/fan-zhuan-lian-biao-shuang-zhi-zhen-di-gui-yao-mo-/
func reverseList(head *ListNode) *ListNode {
	var curr *ListNode
	prev := head

	for prev != nil {
		temp := prev.Next
		prev.Next = curr
		curr = prev
		prev = temp
	}
	return curr
}
