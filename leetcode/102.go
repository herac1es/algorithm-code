package leetcode

import (
	"container/list"
)

func levelOrder(root *TreeNode) [][]int {

	q1 := list.New()
	size := 0
	nextSize := 0
	if root != nil {
		q1.PushBack(root)
		size++
	}
	item := make([]int, 0)
	ret := make([][]int, 0)
	for q1.Len() > 0 {
		front := q1.Front().Value.(*TreeNode)
		item = append(item, front.Val)
		if front.Left != nil {
			q1.PushBack(front.Left)
			nextSize++
		}
		if front.Right != nil {
			q1.PushBack(front.Right)
			nextSize++
		}
		q1.Remove(q1.Front())
		size--
		if size == 0 {
			size = nextSize
			nextSize = 0
			ret = append(ret, item)
			item = make([]int, 0)
		}
	}
	return ret
}
