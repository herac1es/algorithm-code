package main

import (
	"fmt"
)

func main() {
	inter := [][]int{{-10, 0}, {1, 5}, {7, 10}, {15, 20}}
	newInter := []int{2, 8}
	fmt.Println(InsertIntervals(inter, newInter))

}

func InsertIntervals(intervals [][]int, newInterval []int) [][]int {

	ret := make([][]int, 0)
	n := len(intervals)
	if n == 0 {
		ret = append(ret, newInterval)
		return ret
	}
	i := 0
	for ; i < n; i++ {

		if !(intervals[i][1] < newInterval[0] || intervals[i][0] > newInterval[1]) {
			break
		}
		ret = append(ret, intervals[i])
	}

	newGap := make([]int, 2)
	newGap[0] = min(intervals[i][0], newInterval[0])
	newGap[1] = max(intervals[i][1], newInterval[1])

	ret = append(ret, newGap)
	i++
	for i < n {
		if intervals[i][1] < newInterval[0] || intervals[i][0] > newInterval[1] {
			ret = append(ret, intervals[i:]...)
			break
		}
		last := ret[len(ret)-1]
		last[0] = min(intervals[i][0], last[0])
		last[1] = max(intervals[i][1], last[1])
		i++
	}
	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
