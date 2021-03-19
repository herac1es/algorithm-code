package classic

// 选择排序: 选择的意思是只每一趟内循环选择一个最大（或者最小的值）
// 交换次数是N次
// 时间：O(n2)
// 空间：O(1)
func SelectionSort(nums []int) []int {
	min := 0
	for i := 0; i < len(nums); i++ {
		min = i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
	return nums
}
