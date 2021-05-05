package jzoffer

// 时间: O(m+n)
// 空间: O(m+n)
func isMatch(s string, p string) bool {
	if p == "" {
		return s == ""
	}
	first := len(s) > 0 && (s[0] == p[0] || p[0] == '.')
	if len(p) >= 2 && p[1] == '*' {
		// 匹配0个，跳过a* || 匹配一个，继续向后匹配
		return isMatch(s, p[2:]) || (first && isMatch(s[1:], p))
	}
	return first && isMatch(s[1:], p[1:])
}
