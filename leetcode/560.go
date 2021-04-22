package leetcode

// 560.和为k的子数组

// 前缀和preffix[i] : 从nums[0]到sums[i]的累和
// nums[i] = preffix[i]-preffix[i-1]
// 当i==0,令 preffix[-1] = 0
// nums[i]+...+nums[j] = preffix[j] - preffix[i-1]
func subarraySum(nums []int, k int) int {
	preffix := make(map[int]int, len(nums)) // key:前缀和 val:出现的次数
	preffix[0] = 1                          // 将preffix[-1] = 0设置
	pref := 0
	ret := 0
	for _, v := range nums {
		pref += v
		if v, ok := preffix[pref-k]; ok {
			ret += v
		}
		preffix[pref]++
	}
	return ret
}

// 暴力枚举
// 时间:O(n2)
// 空间:O(1)
// func subarraySum(nums []int, k int) int {
//     ret := 0

//     for i := 0; i < len(nums); i++ {

//         sum := 0
//         for j := i; j < len(nums); j++ {
//             sum += nums[j]
//             if sum == k {
//                 ret++
//             }
//         }
//     }
//     return ret
// }
