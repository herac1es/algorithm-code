package leetcode

// 46.全排列
// 时间：O(nxn!)
// 空间: O(n)
func permute(nums []int) [][]int {

	ret := make([][]int, 0)
	visited := make(map[int]struct{}, len(nums))

	permuteRecur(nums, visited, nil, &ret)
	return ret
}

func permuteRecur(nums []int, visited map[int]struct{}, item []int, ret *[][]int) {
	if len(visited) == len(nums) {
		*ret = append(*ret, item)
		return
	}
	for _, v := range nums {
		if _, ok := visited[v]; ok {
			continue
		}
		visited[v] = struct{}{}
		permuteRecur(nums, visited, append(item, v), ret)
		delete(visited, v)
	}
}
