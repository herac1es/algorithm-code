package leetcode

// 给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。
//
// 字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。
//
// 说明：
//
//
// 字母异位词指字母相同，但排列不同的字符串。
// 不考虑答案输出的顺序。
//
//
// 示例 1:
//
//
// 输入:
// s: "cbaebabacd" p: "abc"
//
// 输出:
// [0, 6]
//
// 解释:
// 起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
// 起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词。
//
//
// 示例 2:
//
//
// 输入:
// s: "abab" p: "ab"
//
// 输出:
// [0, 1, 2]
//
// 解释:
// 起始索引等于 0 的子串是 "ab", 它是 "ab" 的字母异位词。
// 起始索引等于 1 的子串是 "ba", 它是 "ab" 的字母异位词。
// 起始索引等于 2 的子串是 "ab", 它是 "ab" 的字母异位词。
//
// Related Topics 哈希表
// 👍 491 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
// 滑动窗口
// 时间: O(n)
// 空间: O(1)
func findAnagrams(s string, p string) []int {
	ret := make([]int, 0)
	std := [26]int{0}
	for i := range p {
		std[p[i]-97]++
	}
	nums := [26]int{0}
	l, r := 0, -1
	for l < len(s) {
		if r+1 < len(s) && nums[s[r+1]-97] < std[s[r+1]-97] {
			r++
			nums[s[r]-97]++
		} else {
			nums[s[l]-97]--
			l++
		}
		if r-l+1 == len(p) {
			ret = append(ret, l)
		}
	}
	return ret
}

// leetcode submit region end(Prohibit modification and deletion)

// 暴力滑窗
// 时间: O(n2)
// 空间: O(1)
func findAnagramsII(s string, p string) []int {
	if p == "" || s == "" {
		return []int{}
	}
	ret := make([]int, 0)
	std := [26]int{0}
	for i := range p {
		std[p[i]-97]++
	}
	for i := 0; i <= len(s)-len(p); i++ {
		nums := [26]int{0}
		for j := i; j <= i+len(p)-1; j++ {
			nums[s[j]-97]++
		}
		valid := true
		for k := range p {
			if nums[p[k]-97] == std[p[k]-97] {
				continue
			} else {
				valid = false
				break
			}
		}
		if valid {
			ret = append(ret, i)
		}
	}
	return ret
}
