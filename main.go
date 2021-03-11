package main

import "github.com/iwillbesober/algorithm-code/leetcode"

// 测试用例:["LRUCache","put","put","put","put","get","get","get","get","put","get","get","get","get","get"]
//				[[3],[1,1],[2,2],[3,3],[4,4],[4],[3],[2],[1],[5,5],[1],[2],[3],[4],[5]]
//				测试结果:[null,null,null,null,null,4,3,2,-1,null,-1,2,-1,4,5]
//				期望结果:[null,null,null,null,null,4,3,2,-1,null,-1,2,3,-1,5]
//				stdout:
func main() {
	type c struct {
		arg1 []int
		arg2 []int
		want []int
	}
	cases := []c{
		{
			arg1: []int{4, 1, 2},
			arg2: []int{1, 3, 4, 2},
			want: []int{-1, 3, -1},
		},
		{
			arg1: []int{2, 4},
			arg2: []int{1, 2, 3, 4},
			want: []int{3, -1},
		},
		{
			arg1: []int{1},
			arg2: []int{1},
			want: []int{-1},
		},
	}
	for i := range cases{
		got := leetcode.NextGreaterElement(cases[i].arg1,cases[i].arg2)
		if
	}
}
