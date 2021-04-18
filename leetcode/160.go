package leetcode

// 160.相交链表
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 时间:O(m+n)
// 空间:O(1)
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	exch := 0
	a, b := headA, headB
	for a != nil && b != nil {
		if a == b {
			return a
		}
		if a.Next == nil && exch < 2 {
			a = headB
			exch++
		} else {
			a = a.Next
		}
		if b.Next == nil && exch < 2 {
			b = headA
			exch++
		} else {
			b = b.Next
		}
	}
	return nil
}
