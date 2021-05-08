package jzoffer

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor37() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	res := make([]string, 0)
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			res = append(res, "null")
		} else {
			res = append(res, strconv.Itoa(node.Val))
			queue = append(queue, node.Left, node.Right)
		}
	}
	ret := ""
	if len(res) > 0 {
		ret += "[" + strings.Join(res, ",") + "]"
	} else {
		ret += "[]"
	}
	return ret
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "[]" {
		return nil
	}
	data = data[1 : len(data)-1]
	nodes := strings.Split(data, ",")
	v, _ := strconv.Atoi(nodes[0])
	root := &TreeNode{
		Val: v,
	}
	queue := make([]*TreeNode, 1)
	queue[0] = root
	cursor := 1
	for cursor < len(nodes) {
		n := queue[0]
		queue = queue[1:]
		if nodes[cursor] != "null" {
			v, _ := strconv.Atoi(nodes[cursor])
			left := &TreeNode{
				Val: v,
			}
			n.Left = left
			queue = append(queue, left)
		}
		if nodes[cursor+1] != "null" {
			v, _ := strconv.Atoi(nodes[cursor+1])
			right := &TreeNode{
				Val: v,
			}
			n.Right = right
			queue = append(queue, right)
		}
		cursor += 2
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
