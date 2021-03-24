package leetcode

// 有效的字母异位词
// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
// 时间:O(m+n)
// 空间:O(1)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := make(map[uint8]int, 26)
	for i := range s {
		count[s[i]]++
	}
	for i := range t {
		v := count[t[i]]
		if v == 0 {
			return false
		}
		if v == 1 {
			delete(count, t[i])
		} else {
			count[t[i]] = v - 1
		}
	}
	return len(count) == 0
}
