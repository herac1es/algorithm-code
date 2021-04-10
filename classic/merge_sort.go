package classic

// 归并排序
// 时间：O(nlgN)
// 空间：O(n)
var aux []int

func MergeSort(nums []int) []int {
	// mergeSort(nums, 0, len(nums)-1)
	mergeSort(nums)
	return nums
}

// 自底向上
func mergeSort(nums []int) {
	n := len(nums)
	for size := 1; size < n; size *= 2 {
		for i := 0; i < n-size; i += size * 2 {
			merge(nums, i, i+size-1, min(i+size*2-1, n-1))
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 自顶向下
func mergeSortII(nums []int, l, r int) {
	if l >= r {
		return
	}
	mid := l + (r-l)/2
	mergeSortII(nums, l, mid)
	mergeSortII(nums, mid+1, r)
	merge(nums, l, mid, r)
}

func merge(nums []int, l, mid, r int) {
	if len(aux) == 0 {
		aux = make([]int, len(nums))
	}
	i, j := l, mid+1 // 对[l:mid-1]和[mid:r]归并
	aux = append(aux[:0], nums...)
	for k := l; k <= r; k++ {
		if i > mid {
			nums[k] = aux[j]
			j++
		} else if j > r {
			nums[k] = aux[i]
			i++
		} else if aux[i] <= aux[j] {
			nums[k] = aux[i]
			i++
		} else {
			nums[k] = aux[j]
			j++
		}
	}
}
