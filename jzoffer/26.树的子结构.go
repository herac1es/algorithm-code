package jzoffer

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 时间: O(mn) m，n分别为A，B的节点个数
// 空间: O(m)
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	// 找到一个相等的点，开始递归是否包含B
	// 若不满足，则继续下一个点
	if A.Val == B.Val && isSubTree(A, B) {
		return true
	}
	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func isSubTree(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	if A.Val != B.Val {
		return false
	}
	return isSubTree(A.Left, B.Left) && isSubTree(A.Right, B.Right)
}
