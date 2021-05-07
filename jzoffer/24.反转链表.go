package jzoffer

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 头插法
// 时间: O（n）
// 空间： O(1)
func reverseList(head *ListNode) *ListNode {
	dummy := new(ListNode)
	cur := head
	next := (*ListNode)(nil)
	for cur != nil {
		next = cur.Next
		cur.Next = dummy.Next
		dummy.Next = cur
		cur = next
	}
	return dummy.Next
}
