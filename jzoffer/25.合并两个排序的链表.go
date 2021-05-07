package jzoffer

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 时间 ： O(m+n)
// 空间 : O(1)
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	ptr := dummy
	for l1 != nil || l2 != nil {
		if l1 == nil {
			ptr.Next = l2
			break
		} else if l2 == nil {
			ptr.Next = l1
			break
		} else if l1.Val <= l2.Val {
			ptr.Next = l1
			l1 = l1.Next
		} else {
			ptr.Next = l2
			l2 = l2.Next
		}
		ptr = ptr.Next
	}
	return dummy.Next
}
