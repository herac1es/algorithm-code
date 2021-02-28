package main

import (
	"fmt"
	"sort"
)

func main() {
	inters := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	sort.Slice(inters, func(i, j int) bool {
		return inters[i][1] < inters[j][1]
	})
	x := inters[0]
	res := 1
	for i := 1; i < len(inters); i++ {
		if inters[i][0] >= x[1] {
			x = inters[i]
			res++
		}
	}
	fmt.Println(len(inters) - res)
}
