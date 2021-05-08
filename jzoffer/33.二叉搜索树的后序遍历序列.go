package jzoffer

// 后序遍历：左子树｜右子树｜根节点
// 时间: O(n2)
// 空间: O(height)
func verifyPostorder(postorder []int) bool {
	return verifyRecur(postorder, 0, len(postorder)-1)
}

func verifyRecur(postorder []int, left, right int) bool {
	if left >= right {
		return true
	}
	i := left
	for postorder[i] < postorder[right] { // 找到第一个大于根节点的数
		i++
	}
	// 遍历右子树，检查是否都大于根节点
	j := i
	for j < right && postorder[j] > postorder[right] {
		j++
	}
	return j == right && verifyRecur(postorder, left, i-1) && verifyRecur(postorder, i, right-1)
}
