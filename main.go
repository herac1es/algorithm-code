package main

func main() {

}

func findCombination(nums []int, k int) int {
	cnt := make(map[int]int, len(nums))
	ret := 0
	for _, v := range nums {
		if count, ok := cnt[k-v]; ok {
			ret += count
		}
		cnt[v]++
	}
	return ret
}

func quickSort(nums []int, l, r int) {

	if l >= r {
		return
	}

	mid := partition(nums, l, r)
	quickSort(nums, l, mid-1)
	quickSort(nums, mid+1, r)

}

func partition(nums []int, lo, hi int) int {
	l := lo + 1
	r := hi
	mid := nums[lo]
	for {
		if l < hi && nums[l] <= mid {
			l++
		}
		if r > lo && nums[r] >= mid {
			r--
		}
		if l >= r {
			break
		}
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
	nums[lo], nums[r] = nums[r], nums[lo]
	return r
}
