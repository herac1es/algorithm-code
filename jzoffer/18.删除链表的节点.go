package jzoffer

// 时间: O(n)
// 空间: O(1)
func deleteNode(head *ListNode, val int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	ptr := dummy
	for ptr != nil {
		if ptr.Next != nil && ptr.Next.Val == val {
			ptr.Next = ptr.Next.Next
			break
		}
		ptr = ptr.Next
	}
	return dummy.Next
}
