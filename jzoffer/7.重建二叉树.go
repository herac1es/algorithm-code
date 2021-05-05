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
// 空间: O(n)
func buildTree(preorder []int, inorder []int) *TreeNode {
	ht := make(map[int]int, len(inorder))
	for i, v := range inorder {
		ht[v] = i
	}
	return build(preorder, ht, 0, 0, len(inorder)-1)
}

// 前序遍历： 根节点-左子树-右子树
// 中序遍历： 左子树-根节点-右子树
// 根据前序遍历的根节点，找到在中序遍历中的位置，将前序遍历分成左右子树，递归建立二叉树
// @param: ht:中序遍历节点与索引的映射 root：根节点在前序遍历的索引
// left： 中序根节点对应的树的范围左边界
// right： 中序根节点对应的树的范围右边界
func build(preorder []int, ht map[int]int, root, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	idx := ht[preorder[root]]
	node := new(TreeNode)
	node.Val = preorder[root]
	node.Left = build(preorder, ht, root+1, left, idx-1)
	// 左子树长度: idx-left
	node.Right = build(preorder, ht, root+1+idx-left, idx+1, right)
	return node
}
