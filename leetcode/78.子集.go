package leetcode

func subsets(nums []int) [][]int {
	ret := make([][]int, 0)
	subsetsBT(nums, -1, nil, &ret)
	return ret
}

func subsetsBT(nums []int, idx int, item []int, ret *[][]int) {

	if len(item) == len(nums) {
		*ret = append(*ret, append([]int{}, item...))
		return
	}

	*ret = append(*ret, append([]int{}, item...))

	for i := idx + 1; i < len(nums); i++ {
		subsetsBT(nums, i, append(item, nums[i]), ret)
	}
}
