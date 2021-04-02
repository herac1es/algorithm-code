package leetcode

//在上次打劫完一条街道之后和一圈房屋后，小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为“根”。 除了“根”之外，每栋房子有且只有一个“父“
//房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。
//
// 计算在不触动警报的情况下，小偷一晚能够盗取的最高金额。
//
// 示例 1:
//
// 输入: [3,2,3,null,3,null,1]
//
//     3
//    / \
//   2   3
//    \   \
//     3   1
//
//输出: 7
//解释: 小偷一晚能够盗取的最高金额 = 3 + 3 + 1 = 7.
//
// 示例 2:
//
// 输入: [3,4,5,1,3,null,1]
//
//     3
//    / \
//   4   5
//  / \   \
// 1   3   1
//
//输出: 9
//解释: 小偷一晚能够盗取的最高金额 = 4 + 5 = 9.
//
// Related Topics 树 深度优先搜索
// 👍 737 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func RobIII(root *TreeNode) int {
	return calRobIII(root, true)
}

// calRobIII 约定：pickMe为false时强制不偷此节点,为true可偷可不偷
func calRobIII(root *TreeNode, pickMe bool) int {
	if root == nil {
		return 0
	}
	if !pickMe {
		ret := calRobIII(root.Left, true) + calRobIII(root.Right, true)
		return ret
	}
	ret := max(root.Val+calRobIII(root.Left, false)+calRobIII(root.Right, false), calRobIII(root.Left, true)+calRobIII(root.Right, true))
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)

// 用空间换时间

var cmap = make(map[*TreeNode]map[bool]int)

func robV2(root *TreeNode) int {
	return calRobIIIV2(root, true)
}

func calRobIIIV2(root *TreeNode, pickMe bool) int {
	if root == nil {
		return 0
	}
	if v, ok := cmap[root][pickMe]; ok {
		return v
	}
	if !pickMe {
		ret := calRobIIIV2(root.Left, true) + calRobIIIV2(root.Right, true)
		if cmap[root] == nil {
			cmap[root] = map[bool]int{}
		}
		cmap[root][pickMe] = ret
		return ret
	}
	ret := max(root.Val+calRobIIIV2(root.Left, false)+calRobIIIV2(root.Right, false), calRobIIIV2(root.Left, true)+calRobIIIV2(root.Right, true))
	if cmap[root] == nil {
		cmap[root] = map[bool]int{}
	}
	cmap[root][pickMe] = ret
	return ret
}
