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
// 空间：O(height) ,最差O（n）
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricRecur(root.Left, root.Right)
}

func isSymmetricRecur(A *TreeNode, B *TreeNode) bool {
	if A == nil && B == nil {
		return true
	}
	if A == nil || B == nil {
		return false
	}
	if A.Val != B.Val {
		return false
	}
	return isSymmetricRecur(A.Left, B.Right) && isSymmetricRecur(A.Right, B.Left)
}
