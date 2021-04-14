package leetcode

// 对称二叉树
// 时间：O(n)
// 空间: O(height)
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return isSymmetricRecur(root.Left, root.Right)
}

func isSymmetricRecur(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return isSymmetricRecur(left.Left, right.Right) && isSymmetricRecur(left.Right, right.Left)
}
