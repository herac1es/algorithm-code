package jzoffer

// 以最右元素作为参考点，判断mid所在位置，再不断缩小区间，直到区间里只剩一个数
// 时间: O(logn)
// 空间: O(1)
func minArray(numbers []int) int {
	i, j := 0, len(numbers)-1
	for i < j {
		mid := i + (j-i)/2
		if numbers[mid] > numbers[j] {
			i = mid + 1
		} else if numbers[mid] < numbers[j] {
			j = mid
		} else {
			j--
		}
	}
	return numbers[i]
}
