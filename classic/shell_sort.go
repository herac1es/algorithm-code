package classic

import "fmt"

// 希尔排序
func ShellSort(nums []int) []int {
	fmt.Println(nums)
	n := len(nums)
	w := 1
	for w < n/3 {
		w = 3*w + 1
	}
	for w >= 1 {
		for i := w; i < n; i++ {
			fmt.Println("w: ", i)
			for j := i; j-w >= 0 && nums[j-w] > nums[j]; j -= w {
				fmt.Println(j-w, j)
				nums[j-w], nums[j] = nums[j], nums[j-w]
			}
		}
		w /= 3
	}
	return nums
}
