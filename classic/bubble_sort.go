package classic

// 时间:O(n^2)
// 空间：O(1)
func BubbleSort(nums []int) []int {
	for j := len(nums) - 1; j >= 0; j-- {
		for i := 0; i < j; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}
	return nums
}
