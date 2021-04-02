package leetcode

//给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
//
// 你可以假设除了数字 0 之外，这两个数字都不会以零开头。
//
//
//
// 进阶：
//
// 如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。
//
//
//
// 示例：
//
// 输入：(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 8 -> 0 -> 7
//
// Related Topics 链表
// 👍 362 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbersII(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := make([]uint8, 2)
	s2 := make([]uint8, 2)
	for l1 != nil {
		s1 = append(s1, uint8(l1.Val))
		l1 = l1.Next
	}
	for l2 != nil {
		s2 = append(s2, uint8(l2.Val))
		l2 = l2.Next
	}
	longer, shorter := s1, s2
	if len(s2) > len(s1) {
		longer, shorter = shorter, longer
	}
	j := len(longer) - 1
	i := len(shorter) - 1
	surplus := uint8(0)
	for {
		if surplus == 0 && j < 0 {
			break
		}
		if j >= 0 {
			surplus += longer[j]
		}
		if i >= 0 {
			surplus += shorter[i]
		}
		longer[j] = surplus % 10
		surplus /= 10
		j--
		i--
	}
	ret := new(ListNode)
	cur := ret
	i = 0
	for i < 2 && longer[i] == 0 {
		i++
	}
	for ; i < len(longer); i++ {
		cur.Next = &ListNode{Val: int(longer[i])}
		cur = cur.Next
	}
	return ret.Next
}

//leetcode submit region end(Prohibit modification and deletion)
