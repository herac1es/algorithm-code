package leetcode

import "sort"

func threeSumClosest(nums []int, target int) int {
	ret := 0
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 0; i <= len(nums)-3; i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[l] + nums[r] + nums[i]
			if abs(sum, target) < abs(ret, target) || (i == 0 && l == 1 && r == len(nums)-1) {
				ret = sum
			}
			if sum == target {
				return sum
			} else if sum > target {
				r--
			} else {
				l++
			}
		}
	}
	return ret
}

func abs(a, b int) int {
	ret := a - b
	if ret >= 0 {
		return ret
	}
	return -ret
}
