package leetcode

//请判断一个链表是否为回文链表。
//
// 示例 1:
//
// 输入: 1->2
//输出: false
//
// 示例 2:
//
// 输入: 1->2->2->1
//输出: true
//
//
// 进阶：
//你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
// Related Topics 链表 双指针
// 👍 937 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	ret := true
	n := 0
	cur := head
	for cur != nil {
		cur = cur.Next
		n++
	}
	if n == 1 {
		return true
	}
	var last *ListNode
	cur = head
	for i := 0; i < n/2; i++ {
		if i == n/2-1 {
			last = cur
		}
		cur = cur.Next
	}

	if n%2 != 0 {
		last = cur
		cur = cur.Next
	}
	cur = reverseListDD(cur)
	tail := cur
	pre := head
	for i := 0; i < n/2; i++ {
		if tail.Val != pre.Val {
			ret = false
			break
		}
		tail = tail.Next
		pre = pre.Next
	}
	// 还原链表
	cur = reverseListDD(cur)
	last.Next = cur
	printList(head)
	return ret
}

// 时间 O(n)
// 空间 O(n)
func isPalindromeOn(head *ListNode) bool {
	num := make([]int, 0)
	for head != nil {
		num = append(num, head.Val)
		head = head.Next
	}
	l, r := 0, len(num)-1
	for l < r {
		if num[l] != num[r] {
			return false
		}
		l++
		r--
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
