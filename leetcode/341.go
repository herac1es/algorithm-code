package leetcode

//给你一个嵌套的整型列表。请你设计一个迭代器，使其能够遍历这个整型列表中的所有整数。
//
// 列表中的每一项或者为一个整数，或者是另一个列表。其中列表的元素也可能是整数或是其他列表。
//
//
//
// 示例 1:
//
// 输入: [[1,1],2,[1,1]]
//输出: [1,1,2,1,1]
//解释: 通过重复调用 next 直到 hasNext 返回 false，next 返回的元素的顺序应该是: [1,1,2,1,1]。
//
// 示例 2:
//
// 输入: [1,[4,[6]]]
//输出: [1,4,6]
//解释: 通过重复调用 next 直到 hasNext 返回 false，next 返回的元素的顺序应该是: [1,4,6]。
//
// Related Topics 栈 设计
// 👍 321 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * // this is the interface that allows for creating nested lists.
 * // you should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedIterator struct {
	stack []int
}

func Constructor2(nestedList []*NestedInteger) *NestedIterator {
	iterator := new(NestedIterator)
	for _, v := range nestedList {
		iterator.iterate(v)
	}
	return iterator
}

func (this *NestedIterator) iterate(ni *NestedInteger) {
	if ni.IsInteger() {
		this.stack = append(this.stack, ni.GetInteger())
		return
	}
	for _, v := range ni.GetList() {
		this.iterate(v)
	}
}

func (this *NestedIterator) Next() int {
	if !this.HasNext() {
		panic("no more data")
	}
	ret := this.stack[0]
	this.stack = this.stack[1:]
	return ret
}

func (this *NestedIterator) HasNext() bool {
	return len(this.stack) > 0
}

//leetcode submit region end(Prohibit modification and deletion)

// 这是leetcode实现的接口
type NestedInteger struct {
}

func (this NestedInteger) IsInteger() bool { return false }

func (this NestedInteger) GetInteger() int { return 0 }

func (n *NestedInteger) SetInteger(value int)      {}
func (this *NestedInteger) Add(elem NestedInteger) {}
func (this NestedInteger) GetList() []*NestedInteger {
	return nil
}
