package leetcode

//给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
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
// 返回锯齿形层序遍历如下：
//
//
//[
//  [3],
//  [20,9],
//  [15,7]
//]
//
// Related Topics 栈 树 广度优先搜索
// 👍 435 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}

	queque := make([]*TreeNode, 1)
	queque[0] = root
	h := 1 // 记录层数

	for len(queque) > 0 {
		size := len(queque)

		item := make([]int, size)

		for i, v := range queque {
			if v.Left != nil {
				queque = append(queque, v.Left)
			}
			if v.Right != nil {
				queque = append(queque, v.Right)
			}
			idx := i
			if h%2 == 0 {
				idx = size - 1 - i
			}
			item[idx] = v.Val
		}
		queque = queque[size:]

		ret = append(ret, item)
		h++
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
