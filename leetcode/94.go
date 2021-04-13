package leetcode

//给定一个二叉树的根节点 root ，返回它的 中序 遍历。
//
//
//
// 示例 1：
//
//
//输入：root = [1,null,2,3]
//输出：[1,3,2]
//
//
// 示例 2：
//
//
//输入：root = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：root = [1]
//输出：[1]
//
//
// 示例 4：
//
//
//输入：root = [1,2]
//输出：[2,1]
//
//
// 示例 5：
//
//
//输入：root = [1,null,2]
//输出：[1,2]
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [0, 100] 内
// -100 <= Node.val <= 100
//
//
//
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树 哈希表
// 👍 885 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {

	return orderIteration(root, func(root *TreeNode) []Command {
		if root == nil {
			return nil
		}
		return []Command{
			NewCommand(Op.AddFunc, root.Right),
			NewCommand(Op.GetVal, root),
			NewCommand(Op.AddFunc, root.Left),
		}
	})

	// ret := make([]int, 0)
	// traversal(root, &ret)
	// return ret
}

func traversal(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}
	traversal(root.Left, ret)
	*ret = append(*ret, root.Val)
	traversal(root.Right, ret)
}

//leetcode submit region end(Prohibit modification and deletion)
