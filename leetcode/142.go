package leetcode

// 142.环形链表 II
// 哈希表
// 时间：O(n)
// 空间：O(n)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	visited := make(map[*ListNode]struct{}, 0)
	for head != nil {
		if _, ok := visited[head]; ok {
			return head
		}
		visited[head] = struct{}{}
		head = head.Next
	}
	return nil
}
