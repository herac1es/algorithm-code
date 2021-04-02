package leetcode

//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4]
//输出：[2,1,4,3]
//
//
// 示例 2：
//
//
//输入：head = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：head = [1]
//输出：[1]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目在范围 [0, 100] 内
// 0 <= Node.val <= 100
//
//
//
//
// 进阶：你能在不修改链表节点值的情况下解决这个问题吗?（也就是说，仅修改节点本身。）
// Related Topics 递归 链表
// 👍 870 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 递归
// 时间: O(n)
// 空间: O(n)
func swapPairsII(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := swapPairsII(head.Next.Next)
	after := head.Next
	head.Next = next
	after.Next = head
	return after
}

// 迭代
// 时间：O(n)
// 空间: O(1)
func swapPairs(head *ListNode) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	pre := dummy
	cur := head
	for cur != nil && cur.Next != nil {
		next := cur.Next
		pre.Next = next
		cur.Next = next.Next
		next.Next = cur
		pre = cur
		cur = cur.Next
	}
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
