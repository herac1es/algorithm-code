package leetcode

// func majorityElement(nums []int) int {
//     if len(nums) == 0 {
//         return -1
//     }
//     count := make(map[int]int, len(nums))
//     ret := nums[0]
//     for _, num := range nums {
//         count[num]++
//         if count[num] > count[ret] {
//             ret = num
//         }
//     }
//     return ret
// }

// 时间：O(n)
// 空间：O(1)
func majorityElement(nums []int) int {
	candicate, vote := 0, 0
	for _, v := range nums {
		if vote == 0 { // 重新选举
			candicate = v
		}
		if v != candicate {
			vote--
		} else {
			vote++
		}
	}
	return candicate
}
