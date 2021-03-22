package leetcode

// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
//
//
// 示例 1:
//
//
// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
//
// 示例 2:
//
//
// 输入: s = "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
//
// 示例 3:
//
//
// 输入: s = "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//      请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
//
// 示例 4:
//
//
// 输入: s = ""
// 输出: 0
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 5 * 104
// s 由英文字母、数字、符号和空格组成
//
// Related Topics 哈希表 双指针 字符串 Sliding Window
// 👍 5172 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
// 滑动窗口
// 时间：O(n)
// 空间：O(1)
func lengthOfLongestSubstring(s string) int {
	l, r := 0, -1
	ret := 0
	count := [128]int{0}
	for l < len(s) {
		if r+1 < len(s) && count[s[r+1]] == 0 {
			r++
			count[s[r]]++
		} else {
			count[s[l]]--
			l++
		}
		if r-l+1 > ret {
			ret = r - l + 1
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)

// hashmap记录法:
// 时间：O(n)
// 空间：O(n)
// func lengthOfLongestSubstring(s string) int {
// 	Map := make(map[byte]int, len(s))
// 	ret := 0
// 	depth := 0
// 	for i := 0; i < len(s); i++ {
// 		depth++
// 		if lastIdx, ok := Map[s[i]]; ok {
// 			if i-lastIdx < depth {
// 				depth = i - lastIdx
// 			}
// 		}
// 		if depth > ret {
// 			ret = depth
// 		}
// 		Map[s[i]] = i
// 	}
// 	return ret
// }

// 动态规划：时间:O(n2)
// 空间：O(n)
// func lengthOfLongestSubstring(s string) int {
// 	if len(s) == 0 {
// 		return 0
// 	}
// 	dp := make([]int, len(s)) // 代码以s[i]结尾的最长不重复子串长度
// 	for i := range dp {
// 		dp[i] = 1
// 	}
// 	for i := 1; i < len(s); i++ {
// 		dp[i] = dp[i-1] + 1
// 		for j := i - 1; j >= 0 && j < len(s) && j >= i-1-dp[i-1]+1; j-- {
// 			if s[j] == s[i] {
// 				dp[i] = i - j
// 				break
// 			}
// 		}
// 	}
// 	ret := 0
// 	for i := range dp {
// 		if dp[i] > ret {
// 			ret = dp[i]
// 		}
// 	}
// 	return ret
// }
