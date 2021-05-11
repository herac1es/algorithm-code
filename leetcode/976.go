package leetcode

//给定由一些正数（代表长度）组成的数组 A，返回由其中三个长度组成的、面积不为零的三角形的最大周长。
//
// 如果不能形成任何面积不为零的三角形，返回 0。
//
//
//
//
//
//
// 示例 1：
//
// 输入：[2,1,2]
//输出：5
//
//
// 示例 2：
//
// 输入：[1,2,1]
//输出：0
//
//
// 示例 3：
//
// 输入：[3,2,3,4]
//输出：10
//
//
// 示例 4：
//
// 输入：[3,6,2,3]
//输出：8
//
//
//
//
// 提示：
//
//
// 3 <= A.length <= 10000
// 1 <= A[i] <= 10^6
//
// Related Topics 排序 数学
// 👍 139 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func largestPerimeter(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	QuickSort(nums, 0, len(nums)-1)
	ret := 0
	for r := len(nums) - 1; r >= 2; r-- {

		if nums[r-1]+nums[r-2] > nums[r] {
			ret = nums[r-1] + nums[r-2] + nums[r]
			break
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)

func QuickSort(nums []int, l, r int) {
	if l >= r {
		return
	}

	j := partition976(nums, l, r)
	QuickSort(nums, l, j-1)
	QuickSort(nums, j+1, r)
}

func partition976(nums []int, l, r int) int {
	i := l + 1
	j := r
	mid := l + (r-l)/2
	nums[mid], nums[l] = nums[l], nums[mid]
	v := nums[l]
	for {
		for i < r && nums[i] < v {
			i++
		}
		for nums[j] > v {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	nums[l], nums[j] = nums[j], nums[l]
	return j
}
