package jzoffer

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 时间: O(n)
// 空间: O(1)
func reversePrint(head *ListNode) []int {
	stack := make([]int, 0)
	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}
	for i := 0; i < len(stack)/2; i++ {
		j := len(stack) - i - 1
		stack[i], stack[j] = stack[j], stack[i]
	}
	return stack
}
