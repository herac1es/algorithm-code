package jzoffer

// 时间: O(n)
// 空间: O(1)
func replaceSpace(s string) string {
	ret := ""
	for _, v := range s {
		if v == ' ' {
			ret += "%20"
		} else {
			ret += string(v)
		}
	}
	return ret
}
