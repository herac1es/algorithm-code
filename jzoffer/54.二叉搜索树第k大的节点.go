package jzoffer

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	ret := 0
	kthRecur(root, &k, &ret)
	return ret
}

// 二叉搜索树中序遍历是递增序列：left->root>right
// --> 二叉树中序遍历取反： right->root->left
// 时间：O(n)
// 空间: O(n) 递归栈的深度
func kthRecur(root *TreeNode, k *int, res *int) {
	if root == nil {
		return
	}
	kthRecur(root.Right, k, res)
	if *k == 0 {
		return
	}
	*k--
	if *k == 0 {
		*res = root.Val
		return
	}
	kthRecur(root.Left, k, res)
}
