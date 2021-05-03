package jzoffer

// 时间: O(n)
// 空间： O(1)
func firstUniqChar(s string) byte {
	hm := [26]int{}
	for _, v := range s {
		hm[v-'a']++
	}
	for _, v := range s {
		if hm[v-'a'] == 1 {
			return byte(v)
		}
	}
	return ' '
}
