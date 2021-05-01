package jzoffer

// 时间：O(n)
// 空间: O(n)
func reverseWords(s string) string {
	if s == "" {
		return ""
	}
	stack := make([]string, 0)
	i, j := 0, -1
	for j < len(s) {

		for i < len(s) && s[i] == ' ' {
			i++
		}
		j = i
		for j < len(s) && s[j] != ' ' {
			j++
		}
		if i < j && j <= len(s) {
			stack = append(stack, s[i:j])
		}
		i = j
	}
	ret := ""
	for len(stack) > 0 {
		ret += stack[len(stack)-1]
		if len(stack) != 1 {
			ret += " "
		}
		stack = stack[:len(stack)-1]
	}
	return ret
}
