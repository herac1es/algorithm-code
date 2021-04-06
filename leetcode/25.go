package leetcode

import "fmt"

//给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
//
// k 是一个正整数，它的值小于或等于链表的长度。
//
// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
// 进阶：
//
//
// 你可以设计一个只使用常数额外空间的算法来解决此问题吗？
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
//
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], k = 2
//输出：[2,1,4,3,5]
//
//
// 示例 2：
//
//
//输入：head = [1,2,3,4,5], k = 3
//输出：[3,2,1,4,5]
//
//
// 示例 3：
//
//
//输入：head = [1,2,3,4,5], k = 1
//输出：[1,2,3,4,5]
//
//
// 示例 4：
//
//
//输入：head = [1], k = 1
//输出：[1]
//
//
//
//
//
// 提示：
//
//
// 列表中节点的数量在范围 sz 内
// 1 <= sz <= 5000
// 0 <= Node.val <= 1000
// 1 <= k <= sz
//
// Related Topics 链表
// 👍 1017 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	dummy := new(ListNode)
	dummy.Next = head
	hd := dummy
	tail := (*ListNode)(nil)
	cur := dummy.Next
	for cur != nil {
		pre := cur
		after := cur.Next
		for i := 0; i < k-1; i++ {
			// 不足k个直接返回
			if cur == nil {
				return dummy.Next
			}
			cur = cur.Next
		}
		if cur == nil {
			return dummy.Next
		}
		hd.Next = cur
		hd = pre
		tail = cur.Next
		cur = tail
		pre.Next = tail

		for j := 0; j < k-1; j++ {
			next := after.Next
			after.Next = pre
			pre = after
			after = next
		}
	}
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d-->", head.Val)
		head = head.Next
	}
	fmt.Println("NULL")
}
