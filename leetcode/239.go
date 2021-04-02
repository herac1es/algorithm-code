package leetcode

import "container/list"

//给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位
//。
//
// 返回滑动窗口中的最大值。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
//输出：[3,3,5,5,6,7]
//解释：
//滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
//
//
// 示例 2：
//
//
//输入：nums = [1], k = 1
//输出：[1]
//
//
// 示例 3：
//
//
//输入：nums = [1,-1], k = 1
//输出：[1,-1]
//
//
// 示例 4：
//
//
//输入：nums = [9,11], k = 2
//输出：[11]
//
//
// 示例 5：
//
//
//输入：nums = [4,-2], k = 2
//输出：[4]
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 105
// -104 <= nums[i] <= 104
// 1 <= k <= nums.length
//
// Related Topics 堆 Sliding Window
// 👍 885 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func maxSlidingWindow(nums []int, k int) []int {
	// if k == 1 {
	// 	return nums
	// }
	res := make([]int, 1)
	window := &MonotonousQueue{
		q: list.New(),
	}
	for i := 0; i < len(nums); i++ {
		if i <= k-1 {
			window.Push(nums[i])
			res[0] = window.Max()
		} else {
			window.Push(nums[i])
			window.Pop(nums[i-k])
			res = append(res, window.Max())
		}
	}
	return res
}

// 使用container/list 实现
type MonotonousQueue struct {
	q *list.List
}

func (m *MonotonousQueue) Push(val int) {
	for m.q.Len() > 0 && m.q.Back().Value.(int) < val {
		m.q.Remove(m.q.Back())
	}
	m.q.PushBack(val)
}

func (m *MonotonousQueue) Pop(val int) {
	if m.q.Len() > 0 && m.q.Front().Value.(int) == val {
		m.q.Remove(m.q.Front())
	}
}

func (m *MonotonousQueue) Max() int {
	return m.q.Front().Value.(int)
}

// 自我实现
// type MonotonousQueue struct {
// 	Back  *DeNode
// 	Front *DeNode
// }
//
// type DeNode struct {
// 	Val       int
// 	Pre, Next *DeNode
// }
//
// func (q *MonotonousQueue) Push(val int) {
// 	for !q.IsEmpty() && q.Back.Val < val {
// 		q.PopBack()
// 	}
// 	q.PushBack(val)
// }
// func (q *MonotonousQueue) Pop(val int) {
// 	if !q.IsEmpty() && q.Front.Val == val {
// 		q.PopFront()
// 	}
// }
//
// func (q *MonotonousQueue) Max() int {
// 	if q.IsEmpty() {
// 		return 0
// 	}
// 	return q.Front.Val
// }
//
// func (q *MonotonousQueue) PopBack() *DeNode {
// 	ret := q.Back
// 	if q.Front == q.Back {
// 		q.Back, q.Front = nil, nil
// 		return ret
// 	}
// 	q.Back = q.Back.Pre
// 	q.Back.Next = nil
// 	return ret
// }
//
// func (q *MonotonousQueue) PushBack(val int) {
// 	node := &DeNode{
// 		Val: val,
// 	}
// 	if q.IsEmpty() {
// 		q.Front = node
// 		q.Back = node
// 		return
// 	}
// 	q.Back.Next = node
// 	node.Pre = q.Back
// 	q.Back = node
// }
//
// func (q *MonotonousQueue) PopFront() *DeNode {
// 	ret := q.Front
// 	if q.Front == q.Back {
// 		q.Front = nil
// 		q.Back = nil
// 	} else {
// 		q.Front = q.Front.Next
// 		q.Front.Pre = nil
// 	}
// 	return ret
// }
//
// func (q *MonotonousQueue) IsEmpty() bool {
// 	return q.Back == q.Front && q.Back == nil
// }

//leetcode submit region end(Prohibit modification and deletion)
