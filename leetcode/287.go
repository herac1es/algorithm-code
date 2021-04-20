package leetcode

// 对于1-n的n个数，找一个中位数mid，然后统计nums[l,r]区间内小于等于他的个数，如果大于n，则重复的数出现在[1-mid]，否则在[mid+1,n]之间
// 时间:O(nlogn)
// 空间:O(1)
func findDuplicate(nums []int) int {
	l, r := 1, len(nums)
	mid, cnt := 0, 0
	for l < r {
		mid = l + (r-l)/2
		cnt = 0
		for _, v := range nums {
			if v <= mid {
				cnt++
			}
		}
		if cnt > mid {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

// func findDuplicate(nums []int) int {
//     hm := make(map[int]struct{},len(nums))
//     for _, v := range nums {
//         if _, ok := hm[v]; ok {
//             return v
//         }
//         hm[v] = struct{}{}
//     }
//     return -1
// }
