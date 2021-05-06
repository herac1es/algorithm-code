package jzoffer

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 时间: O(n)
// 空间: O(1)
func getKthFromEnd(head *ListNode, k int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	pre, after := dummy, dummy
	for i := 0; i < k; i++ {
		pre = pre.Next
		if pre == nil {
			return nil
		}
	}
	for pre != nil {
		pre = pre.Next
		after = after.Next
	}
	return after
}
