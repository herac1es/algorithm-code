package main

import (
	"fmt"

	"github.com/iwillbesober/algorithm-code/leetcode"
)

// 测试用例:["LRUCache","put","put","put","put","get","get","get","get","put","get","get","get","get","get"]
//				[[3],[1,1],[2,2],[3,3],[4,4],[4],[3],[2],[1],[5,5],[1],[2],[3],[4],[5]]
//				测试结果:[null,null,null,null,null,4,3,2,-1,null,-1,2,-1,4,5]
//				期望结果:[null,null,null,null,null,4,3,2,-1,null,-1,2,3,-1,5]
//				stdout:
func main() {
	a := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(leetcode.MaxArea(a))
}
