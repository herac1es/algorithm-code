package jzoffer

type ListNode struct {
	Val  int
	Next *ListNode
}

// 时间: O(a+b)，a,b为两链表长度
// 空间: O(1)
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	ptra := headA
	ptrb := headB

	for ptra != nil && ptrb != nil {
		if ptra == ptrb {
			return ptra
		}
		if ptra.Next == nil && ptrb.Next == nil {
			return nil
		}
		if ptra.Next == nil {
			ptra = headB
		} else {
			ptra = ptra.Next
		}
		if ptrb.Next == nil {
			ptrb = headA
		} else {
			ptrb = ptrb.Next
		}
	}
	return nil
}
