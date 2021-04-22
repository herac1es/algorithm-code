package leetcode

// 543.二叉树的直径

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 543.二叉树的直径
// 时间复杂度:O（n）
// 空间复杂度:O(height)
func diameterOfBinaryTree(root *TreeNode) int {
	ret := 0
	depth(root, &ret)
	return ret
}

func depth(root *TreeNode, ret *int) int {
	if root == nil {
		return 0
	}
	left := depth(root.Left, ret)
	right := depth(root.Right, ret)

	*ret = max(left+right, *ret)
	return max(left, right) + 1
}
