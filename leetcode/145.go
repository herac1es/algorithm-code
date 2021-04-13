package leetcode

//给定一个二叉树，返回它的 后序 遍历。
//
// 示例:
//
// 输入: [1,null,2,3]
//   1
//    \
//     2
//    /
//   3
//
//输出: [3,2,1]
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树
// 👍 564 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	// 迭代
	// return orderIteration(root, func(root *TreeNode) []Command {
	// 	if root == nil {
	// 		return nil
	// 	}
	// 	return []Command{
	// 		NewCommand(Op.GetVal, root),
	// 		NewCommand(Op.AddFunc, root.Right),
	// 		NewCommand(Op.AddFunc, root.Left),
	// 	}
	// })

	ret := make([]int, 0)
	postorderRecursion(root, &ret)
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)

func postorderRecursion(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}
	postorderRecursion(root.Left, ret)
	postorderRecursion(root.Right, ret)
	*ret = append(*ret, root.Val)
}
