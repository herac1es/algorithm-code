package main

import "github.com/iwillbesober/algorithm-code/leetcode"

// 测试用例:["LRUCache","put","put","put","put","get","get","get","get","put","get","get","get","get","get"]
//				[[3],[1,1],[2,2],[3,3],[4,4],[4],[3],[2],[1],[5,5],[1],[2],[3],[4],[5]]
//				测试结果:[null,null,null,null,null,4,3,2,-1,null,-1,2,-1,4,5]
//				期望结果:[null,null,null,null,null,4,3,2,-1,null,-1,2,3,-1,5]
//				stdout:
func main() {
	cache := leetcode.Constructor(3)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Put(4, 4)
	cache.Get(4)
	cache.Get(3)
	cache.Get(2)
	cache.Get(1)
	cache.Put(5, 5)
	cache.Get(1)
	cache.Get(2)
	cache.Get(3)
	cache.Get(4)
	cache.Get(5)
}
