package jzoffer

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 时间：O(n)
// 空间: O(n)
func levelOrderIII(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	level := 1
	for len(queue) > 0 {
		size := len(queue)
		row := make([]int, size)
		for i := 0; i < size; i++ {
			idx := i
			if level%2 == 0 {
				idx = size - 1 - i
			}
			row[idx] = queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		level++
		ret = append(ret, row)
		queue = queue[size:]
	}
	return ret
}
