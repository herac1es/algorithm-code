package jzoffer

// 时间:O(n)
// 空间:O(1)
func isStraight(nums []int) bool {
	quickSort(nums, 0, len(nums)-1)
	zeroCnt := 0
	gapCnt := 0
	for idx := range nums {
		if nums[idx] == 0 {
			zeroCnt++
			continue
		}
		if idx-1 >= 0 && nums[idx-1] != 0 {
			if nums[idx-1] == nums[idx] {
				return false
			}
			gapCnt += nums[idx] - nums[idx-1] - 1
		}
	}
	return gapCnt <= zeroCnt
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}

	j := partition(nums, l, r)
	quickSort(nums, l, j-1)
	quickSort(nums, j+1, r)
}

func partition(nums []int, lo, hi int) int {

	i := lo + 1
	j := hi
	mid := lo + (hi-lo)/2
	nums[mid], nums[lo] = nums[lo], nums[mid]
	v := nums[lo]
	for {
		for i < hi && nums[i] < v {
			i++
		}
		for nums[j] > v {
			j--
		}
		if i >= j { // 相遇，跳出循环
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	nums[lo], nums[j] = nums[j], nums[lo]
	return j
}
