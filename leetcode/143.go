package leetcode

//给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
//将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…
//
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
// 示例 1:
//
// 给定链表 1->2->3->4, 重新排列为 1->4->2->3.
//
// 示例 2:
//
// 给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.
// Related Topics 链表
// 👍 561 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	nodes := make([]*ListNode, 0)
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}
	if len(nodes) <= 2 {
		return
	}
	for i := 0; i < len(nodes)/2; i++ {
		rIdx := len(nodes) - 1 - i
		nodes[i].Next = nodes[rIdx]
		if rIdx != i+1 && i+1 < len(nodes) {
			nodes[rIdx].Next = nodes[i+1]
		}
	}
	nodes[len(nodes)/2].Next = nil
}

//leetcode submit region end(Prohibit modification and deletion)
