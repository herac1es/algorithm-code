package leetcode

//反转一个单链表。
//
// 示例:
//
// 输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL
//
// 进阶:
//你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
// Related Topics 链表
// 👍 1570 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 递归
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	last := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

//leetcode submit region end(Prohibit modification and deletion)

// 迭代法
// 时间: O(n)
// 空间: O(1)
func reverseListDD(head *ListNode) *ListNode {
	pre := (*ListNode)(nil)
	cur := head
	var (
		next *ListNode
	)
	if cur != nil {
		next = cur.Next
	}
	for cur != nil {
		cur.Next = pre
		pre = cur
		cur = next
		if next != nil {
			next = next.Next
		}
	}
	return pre
}
