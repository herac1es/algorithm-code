package jzoffer

// 滑动窗口
// 时间:O(n)
// 空间:O(1)
func lengthOfLongestSubstring(s string) int {
	cnt := [256]int{}
	l, r := 0, -1
	ret := 0
	for l < len(s) && r < len(s) {
		if r-l+1 > ret {
			ret = r - l + 1
		}
		if r+1 < len(s) && cnt[s[r+1]] == 0 {
			r++
			cnt[s[r]]++
		} else {
			cnt[s[l]]--
			l++
		}
	}
	return ret
}
