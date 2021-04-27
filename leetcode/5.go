package leetcode

//给你一个字符串 s，找到 s 中最长的回文子串。
//
//
//
// 示例 1：
//
//
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
//
//
// 示例 2：
//
//
//输入：s = "cbbd"
//输出："bb"
//
//
// 示例 3：
//
//
//输入：s = "a"
//输出："a"
//
//
// 示例 4：
//
//
//输入：s = "ac"
//输出："a"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 仅由数字和英文字母（大写和/或小写）组成
//
// Related Topics 字符串 动态规划
// 👍 3288 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
// 动态规划
// 时间：O(n2)
// 空间: O(n2)
func LongestPalindrome(s string) string {
	n := len(s)
	ret := string(s[0])
	// dp[i][j] : 字符串s[i:j]是否是回文字符串
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	// basecase：每一个字符本身是回文字符
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}
	// dp[i][j]是回文字符串
	// 1: 若 j = i+1，即只有两个字符，则看两个字符是否相等
	// 2: 否则当且仅当 dp[i+1][j-1]为true ，且 s[i] == s[j]
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if j == i+1 && s[i] == s[j] {
				dp[i][j] = true
			} else if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			}
			if dp[i][j] == true && (j-i+1) > len(ret) {
				ret = s[i : j+1]
			}
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
