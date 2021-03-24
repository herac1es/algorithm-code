package classic

// 希尔排序
func ShellSort(nums []int) []int {
	n := len(nums)
	w := 1
	for w < n/3 {
		w = 3*w + 1
	}
	for w >= 1 {
		for i := w; i < n; i++ {
			for j := i; j-w >= 0 && nums[j-w] > nums[j]; j -= w {
				nums[j-w], nums[j] = nums[j], nums[j-w]
			}
		}
		w /= 3
	}
	return nums
}
