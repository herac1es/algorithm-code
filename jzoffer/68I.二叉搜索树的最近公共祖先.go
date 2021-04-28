package jzoffer

// 递归：
// 时间: O(n)
// 空间: O(height)
func lowestCommonAncestorBST(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestorBST(root.Right, p, q)
	}
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestorBST(root.Left, p, q)
	}
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
// 迭代
// 时间：O(n)
// 空间: O(1)
func lowestCommonAncestorII(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}
	ptr := root
	for ptr != nil {
		if ptr.Val < p.Val && ptr.Val < q.Val {
			ptr = ptr.Right
		} else if ptr.Val > p.Val && ptr.Val > q.Val {
			ptr = ptr.Left
		} else {
			return ptr
		}
	}
	return nil
}
