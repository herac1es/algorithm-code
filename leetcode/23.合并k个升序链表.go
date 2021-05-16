package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 时间: O(knxlogk)
// 空间: O(logk)
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	mergedList := make([]*ListNode, 0, len(lists)/2)
	for i := 0; i < len(lists); i += 2 {
		if i+1 < len(lists) {
			mergedList = append(mergedList, merge23(lists[i], lists[i+1]))
		} else {
			mergedList = append(mergedList, lists[i])
		}
	}
	return mergeKLists(mergedList)
}

func merge23(l1 *ListNode, l2 *ListNode) *ListNode {

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
