package leetcode

// 给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
//
// 进阶：
//
//
// 你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
//
//
//
//
// 示例 1：
//
//
// 输入：head = [4,2,1,3]
// 输出：[1,2,3,4]
//
//
// 示例 2：
//
//
// 输入：head = [-1,5,3,4,0]
// 输出：[-1,0,3,4,5]
//
//
// 示例 3：
//
//
// 输入：head = []
// 输出：[]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目在范围 [0, 5 * 104] 内
// -105 <= Node.val <= 105
//
// Related Topics 排序 链表
// 👍 1085 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	n := 0
	for head != nil {
		n++
		head = head.Next
	}
	var (
		pre *ListNode
	)
	for size := 1; size < n; size *= 2 {
		pre = dummy
		cur := dummy.Next
		for cur != nil {
			l1 := &ListNode{}
			lhead := l1

			for i := 0; i < size && cur != nil; i++ {
				l1.Next = cur
				l1 = l1.Next
				cur = cur.Next
			}
			l2 := &ListNode{}
			rhead := l2
			for i := 0; i < size && cur != nil; i++ {
				l2.Next = cur
				cur = cur.Next
				l2 = l2.Next
			}
			head, tail := mergeList(lhead.Next, rhead.Next, l2.Next)
			pre.Next = head
			pre = tail
			cur = pre.Next
		}
	}
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func mergeList(l, r, next *ListNode) (head, tail *ListNode) {
	dummy := new(ListNode)
	cur := dummy
	l2 := r
	for l != l2 || r != next {
		if l == l2 {
			cur.Next = r
			r = r.Next
		} else if r == next {
			cur.Next = l
			l = l.Next
		} else if l.Val < r.Val {
			cur.Next = l
			l = l.Next
		} else {
			cur.Next = r
			r = r.Next
		}
		cur = cur.Next
	}
	cur.Next = next
	return dummy.Next, cur
}
