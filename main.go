package main

import "fmt"

func main() {
	r := new(TreeNode)
	r.Right = &TreeNode{
		Val: 4,
	}
	fmt.Println("hello world")
	kthLargest(r, 2)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	ret := 0
	kthRecur(root, &k, &ret)
	return ret
}

func kthRecur(root *TreeNode, k *int, res *int) {
	if root == nil {
		return
	}
	kthRecur(root.Right, k, res)
	if *k == 0 {
		return
	}
	if root.Left == nil && root.Right == nil {
		*k--
	}
	if *k == 0 {
		*res = root.Val
		return
	}
	kthRecur(root.Left, k, res)
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
