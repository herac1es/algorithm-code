package leetcode

//给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
//
//
//
// 示例 1：
//
//
//输入：root = [1,null,2,3]
//输出：[1,2,3]
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
//输出：[1,2]
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
// 进阶：递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树
// 👍 553 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {

	// 迭代解法
	return preorderIteration(root)

	// 递归解法
	// ret := make([]int, 0)
	// preorderRecursion(root, &ret)
	// return ret
}

//leetcode submit region end(Prohibit modification and deletion)

var Op = struct {
	AddFunc uint8
	GetVal  uint8
}{
	AddFunc: 0,
	GetVal:  1,
}

type Command struct {
	Op   uint8
	Node *TreeNode
}

func NewCommand(op uint8, node *TreeNode) Command {
	return Command{
		Op:   op,
		Node: node,
	}
}

func NewCommandGroup(root *TreeNode) []Command {
	if root == nil {
		return []Command{}
	}
	return []Command{
		NewCommand(Op.AddFunc, root.Right),
		NewCommand(Op.AddFunc, root.Left),
		NewCommand(Op.GetVal, root),
	}
}

// 模拟函数栈
// 时间：O(n)
// 空间: O(n)
func preorderIteration(root *TreeNode) []int {
	ret := make([]int, 0)

	stack := make([]Command, 0, 3)
	stack = append(stack, NewCommandGroup(root)...)

	for len(stack) != 0 {
		c := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if c.Node == nil {
			continue
		}
		if c.Op == Op.AddFunc {
			stack = append(stack, NewCommandGroup(c.Node)...)
		} else {
			ret = append(ret, c.Node.Val)
		}
	}
	return ret
}

// 递归
// 时间: O(n)
// 空间: O(n)
func preorderRecursion(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}
	*ret = append(*ret, root.Val)
	preorderRecursion(root.Left, ret)
	preorderRecursion(root.Right, ret)
}
