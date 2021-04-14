package leetcode

// 二叉树的最大深度
// 递归
// 时间：O(n)
// 空间：O(height) 递归的栈空间
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
