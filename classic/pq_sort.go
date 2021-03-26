package classic

func PqSort(nums []int) []int {
	pqSort(nums)
	return nums
}

func pqSort(nums []int) {
	size := len(nums)
	for i := size / 2; i >= 1; i-- {
		sink(nums, i, size)
	}
	for size > 1 {
		nums[0], nums[size-1] = nums[size-1], nums[0]
		size--
		sink(nums, 1, size)
	}
}

func sink(nums []int, idx int, size int) {
	for (idx << 1) <= size {
		t := idx << 1
		if t+1 <= size && nums[t+1-1] > nums[t-1] {
			t = t + 1
		}
		if nums[t-1] <= nums[idx-1] {
			break
		}
		nums[idx-1], nums[t-1] = nums[t-1], nums[idx-1]
		idx = t
	}
}
