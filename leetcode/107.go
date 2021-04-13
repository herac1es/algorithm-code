package leetcode

//给定一个二叉树，返回其节点值自底向上的层序遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
//
// 例如：
//给定二叉树 [3,9,20,null,null,15,7],
//
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//
//
// 返回其自底向上的层序遍历为：
//
//
//[
//  [15,7],
//  [9,20],
//  [3]
//]
//
// Related Topics 树 广度优先搜索
// 👍 426 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	queque := make([]*TreeNode, 1)
	queque[0] = root
	for len(queque) > 0 {
		size := len(queque)
		res := make([]int, 0, size)

		for _, v := range queque {
			if v.Left != nil {
				queque = append(queque, v.Left)
			}
			if v.Right != nil {
				queque = append(queque, v.Right)
			}
			res = append(res, v.Val)
		}
		ret = append(ret, res)
		queque = queque[size:]
	}
	l, r := 0, len(ret)-1
	for l < r {
		ret[l], ret[r] = ret[r], ret[l]
		l++
		r--
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
