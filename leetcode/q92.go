package leetcode

//反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
//
// 说明:
//1 ≤ m ≤ n ≤ 链表长度。
//
// 示例:
//
// 输入: 1->2->3->4->5->NULL, m = 2, n = 4
//输出: 1->4->3->2->5->NULL
// Related Topics 链表
// 👍 713 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 迭代法
// 时间: O(n)
// 空间: O(1)
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseN(head, right, nil)
	}
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

// 反转前n个节点
func reverseN(head *ListNode, n int, next *ListNode) *ListNode {
	if n == 1 {
		next = head.Next
		return head
	}
	last := reverseN(head.Next, n-1, next)
	head.Next.Next = head
	head.Next = next
	return last
}

//leetcode submit region end(Prohibit modification and deletion)

// 指针法
func reverseBetweenSolutionII(head *ListNode, left int, right int) *ListNode {
	depth := 0
	dummy := new(ListNode)
	dummy.Next = head
	pre := dummy
	cur := new(ListNode)
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	lp := pre // 记录第left节点的前一个节点
	pre = pre.Next
	cur = pre.Next
	var next *ListNode
	if cur != nil {
		next = cur.Next
	}
	depth = 0
	for cur != nil && depth < right-left { // 按照基础反转链表的过程反转从left到right的节点
		cur.Next = pre
		pre = cur
		cur = next
		if next != nil {
			next = next.Next
		}
		depth++
	}
	// 此时pre指向第right个节点
	// lp.Next是反转前的第left个节点
	lp.Next.Next = cur
	lp.Next = pre
	return dummy.Next
}
