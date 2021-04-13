package leetcode

//给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
//
// 示例:
//
// 输入: [1,2,3,null,5,null,4]
//输出: [1, 3, 4]
//解释:
//
//   1            <---
// /   \
//2     3         <---
// \     \
//  5     4       <---
//
// Related Topics 树 深度优先搜索 广度优先搜索 递归 队列
// 👍 446 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {

	ret := make([]int, 0)
	if root == nil {
		return ret
	}

	queue := make([]*TreeNode, 1)
	queue[0] = root

	for len(queue) > 0 {

		size := len(queue)

		for i, v := range queue {
			if v.Left != nil {
				queue = append(queue, v.Left)
			}
			if v.Right != nil {
				queue = append(queue, v.Right)
			}
			if i == size-1 {
				ret = append(ret, v.Val)
			}
		}

		queue = queue[size:]
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
