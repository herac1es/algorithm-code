package jzoffer

// 时间: O(n)
// 空间: O(height)
// 返回值：双向链表的最小节点
func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	// 叶子节点，已自己为前驱和后继
	if root.Left == nil && root.Right == nil {
		root.Left = root
		root.Right = root
		return root
	}
	// l：左子树的循环双向链表最小节点
	// r：右子树的循环双向链表最小节点
	l := treeToDoublyList(root.Left)
	r := treeToDoublyList(root.Right)
	if l != nil && r != nil {
		root.Left = l.Left
		l.Left.Right = root
		root.Right = r
		l.Left = r.Left
		r.Left.Right = l
		r.Left = root
		return l
	} else if l == nil {
		root.Right = r
		root.Left = r.Left
		r.Left.Right = root
		r.Left = root
		return root
	} else {
		root.Left = l.Left
		root.Right = l
		l.Left.Right = root
		l.Left = root
		return l
	}
}
