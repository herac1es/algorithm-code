package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 98.判断二叉搜索树
// 时间：O(n)
// 空间:O(height)
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return true
	} else if root.Left == nil {
		return findMin(root.Right).Val > root.Val && isValidBST(root.Right)
	} else if root.Right == nil {
		return findMax(root.Left).Val < root.Val && isValidBST(root.Left)
	} else {
		return findMax(root.Left).Val < root.Val && findMin(root.Right).Val > root.Val && isValidBST(root.Left) && isValidBST(root.Right)
	}
}

func findMin(root *TreeNode) *TreeNode {
	if root.Left == nil {
		return root
	}
	return findMin(root.Left)
}

func findMax(root *TreeNode) *TreeNode {
	if root.Right == nil {
		return root
	}
	return findMax(root.Right)
}
