package jzoffer

// 时间: O(nlogn)
// 空间: O(n)
func reversePairs(nums []int) int {
	ret := 0
	reserseCnt(nums, 0, len(nums)-1, &ret)
	return ret
}

func reserseCnt(nums []int, l, r int, cnt *int) {
	if l >= r {
		return
	}
	mid := l + (r-l)/2
	reserseCnt(nums, l, mid, cnt)
	reserseCnt(nums, mid+1, r, cnt)
	merge(nums, l, mid+1, r, cnt)
}

func merge(nums []int, l, mid, r int, cnt *int) {
	tmp := make([]int, r-l+1)
	i, j := l, mid
	idx := 0
	for idx < len(tmp) {
		if i >= mid {
			tmp[idx] = nums[j]
			j++
		} else if j > r {
			tmp[idx] = nums[i]
			i++
		} else if nums[i] <= nums[j] {
			tmp[idx] = nums[i]
			i++
		} else { // nums[i] > nums[j] 此时左子数组nums[i:mid)的数和右子数组nums[j]都构成逆序对
			tmp[idx] = nums[j]
			j++
			*cnt += mid - i
		}
		idx++
	}
	idx = l
	for _, v := range tmp {
		nums[idx] = v
		idx++
	}
}
