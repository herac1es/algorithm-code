package leetcode

// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//
// 注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。
//
//
//
// 示例 1：
//
//
// 输入：s = "ADOBECODEBANC", t = "ABC"
// 输出："BANC"
//
//
// 示例 2：
//
//
// 输入：s = "a", t = "a"
// 输出："a"
//
//
//
//
// 提示：
//
//
// 1 <= s.length, t.length <= 105
// s 和 t 由英文字母组成
//
//
//
// 进阶：你能设计一个在 o(n) 时间内解决此问题的算法吗？ Related Topics 哈希表 双指针 字符串 Sliding Window
// 👍 1041 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func minWindow(s string, t string) string {
	want := make(map[uint8]int, len(t))
	for i := range t {
		want[t[i]] += 1
	}
	fullMap := make(map[uint8]struct{})
	count := [128]int{0}
	l, r := 0, -1
	ret := s + "_"
	for l < len(s) {
		if r+1 < len(s) && len(fullMap) < len(want) {
			r++
			count[s[r]]++
			if v, ok := want[s[r]]; ok && count[s[r]] >= v {
				fullMap[s[r]] = struct{}{}
			}
		} else {
			count[s[l]]--
			if v, ok := want[s[l]]; ok && count[s[l]] < v {
				delete(fullMap, s[l])
			}
			l++
		}
		if len(fullMap) == len(want) && r-l+1 < len(ret) {
			ret = s[l : r+1]
		}
	}
	if len(ret) > len(s) {
		ret = ""
	}
	return ret
}

// leetcode submit region end(Prohibit modification and deletion)
