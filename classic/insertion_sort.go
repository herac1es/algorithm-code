package classic

// 插入排序
// 每一步都假定左边的元素已经是有序状态，此时找一个位置将当前元素插入，插入位置右边的所有元素都需要向右移动一位
// 时间复杂度: O(n2)
// 空间复杂度: O(1)
func InsertionSort(nums []int) []int {
	// 按升序排
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0 && nums[j] < nums[j-1]; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
	return nums
}

// TODO 改进
