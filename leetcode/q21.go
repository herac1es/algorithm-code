package leetcode

// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
//
//
// 示例 1：
//
//
// 输入：l1 = [1,2,4], l2 = [1,3,4]
// 输出：[1,1,2,3,4,4]
//
//
// 示例 2：
//
//
// 输入：l1 = [], l2 = []
// 输出：[]
//
//
// 示例 3：
//
//
// 输入：l1 = [], l2 = [0]
// 输出：[0]
//
//
//
//
// 提示：
//
//
// 两个链表的节点数目范围是 [0, 50]
// -100 <= Node.val <= 100
// l1 和 l2 均按 非递减顺序 排列
//
// Related Topics 递归 链表
// 👍 1604 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	mr := new(ListNode)
	p := mr
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	for l1 != nil {
		p.Next = l1
		p = p.Next
		l1 = l1.Next
	}
	for l2 != nil {
		p.Next = l2
		p = p.Next
		l2 = l2.Next
	}
	return mr.Next
}

//leetcode submit region end(Prohibit modification and deletion)
