package jzoffer

import "strconv"

// 时间: O(nlogn)
// 空间: O(n)
func minNumber(nums []int) string {
	quickSort46(nums, 0, len(nums)-1)
	ret := ""
	for _, v := range nums {
		ret += strconv.Itoa(v)
	}
	return ret
}

func quickSort46(nums []int, l, r int) {
	if l >= r {
		return
	}
	j := partition46(nums, l, r)
	quickSort46(nums, l, j-1)
	quickSort46(nums, j+1, r)
}

func less(a, b int) bool {
	stra := strconv.Itoa(a) + strconv.Itoa(b)
	strb := strconv.Itoa(b) + strconv.Itoa(a)
	numa, _ := strconv.Atoi(stra)
	numb, _ := strconv.Atoi(strb)
	return numa < numb
}

func partition46(nums []int, lo, hi int) int {
	i := lo + 1
	j := hi
	mid := lo + (hi-lo)/2
	nums[mid], nums[lo] = nums[lo], nums[mid]
	v := nums[lo]
	for {
		for i < hi && less(nums[i], v) {
			i++
		}
		for less(v, nums[j]) {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	nums[lo], nums[j] = nums[j], nums[lo]
	return j
}
