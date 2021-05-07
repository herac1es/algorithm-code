package jzoffer

// 时间: O(n)
// 空间: O(1)
func validateStackSequences(pushed []int, popped []int) bool {
	size := 0
	for i := 0; i < len(popped); i++ {
		// 栈顶元素不等于当前要出栈的元素，模拟进栈
		for size == 0 || (size <= len(pushed) && pushed[size-1] != popped[i]) {
			size++
		}
		// 当找不到当前需要出栈的元素，说明顺序不合法
		if size > len(pushed) || pushed[size-1] != popped[i] {
			return false
		}
		// 模拟出栈
		pushed = append(pushed[:size-1], pushed[size:]...)
		size--
	}
	return true
}
