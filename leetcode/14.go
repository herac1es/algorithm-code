package leetcode

// 14.最长公共前缀
// 类似于自顶向下归并的思想，分而治之
// 时间: O(mlogn)
// 空间: O(logn) 递归栈的深度
func longestCommonPrefix(strs []string) string {
	return longestCommonPrefixRecur(strs, 0, len(strs)-1)
}

func longestCommonPrefixRecur(strs []string, l, r int) string {
	if l == r {
		return strs[l]
	}

	mid := l + (r-l)/2
	prefixLeft := longestCommonPrefixRecur(strs, l, mid)
	prefixRight := longestCommonPrefixRecur(strs, mid+1, r)
	i := 0
	for ; i < min(len(prefixLeft), len(prefixRight)); i++ {
		if prefixLeft[i] == prefixRight[i] {
			continue
		}
		break
	}
	return prefixLeft[:i]
}
