package jzoffer

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 时间: O(n)
// 空间: O(height)
func pathSum(root *TreeNode, target int) [][]int {
	ret := make([][]int, 0)
	pathRecur(root, target, nil, &ret)
	return ret
}

func pathRecur(root *TreeNode, target int, item []int, ret *[][]int) {
	if root == nil {
		return
	}
	target = target - root.Val
	item = append(item, root.Val)
	// 必须是叶子结点
	if target == 0 && root.Left == nil && root.Right == nil {
		*ret = append(*ret, append([]int{}, item...))
		return
	}
	pathRecur(root.Left, target, item, ret)
	pathRecur(root.Right, target, item, ret)
	item = item[:len(item)-1]
}
