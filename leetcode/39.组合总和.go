package leetcode

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	ret := make([][]int, 0)
	combinationSumRecur(candidates, 0, target, nil, &ret)
	return ret
}

func combinationSumRecur(nums []int, idx, target int, item []int, ret *[][]int) {
	if target == 0 {
		*ret = append(*ret, append([]int{}, item...))
		return
	}
	if idx >= len(nums) || nums[idx] > target {
		return
	}

	for i := idx; i < len(nums); i++ {
		combinationSumRecur(nums, i, target-nums[i], append(item, nums[i]), ret)
	}
}
