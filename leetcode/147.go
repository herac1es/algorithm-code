package leetcode

func insertionSortList(head *ListNode) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	pre := head
	next := &ListNode{}
	cur := &ListNode{}
	for head != nil {
		if head.Val >= pre.Val {
			pre = head
			head = head.Next
			continue
		}
		cur = dummy
		for cur.Next != nil && cur.Next != head && cur.Next.Val < head.Val {
			cur = cur.Next
		}
		next = head.Next
		pre.Next = next
		head.Next = cur.Next
		cur.Next = head
		head = next
	}
	return dummy.Next
}
