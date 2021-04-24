package jzoffer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 使用一个hash表映射 N与N'
// 先遍历一遍链表复制每一个节点
// 再遍历复制后的链表，将random的N更新成N'
// 时间:O(n)
// 空间:O(n)
func copyRandomList(head *Node) *Node {
	dict := make(map[*Node]*Node)
	pHead := head
	clonedDummy := &Node{}
	pClone := clonedDummy
	for pHead != nil {
		clone := &Node{
			Val:    pHead.Val,
			Next:   pHead.Next,
			Random: pHead.Random,
		}
		pClone.Next = clone
		pClone = pClone.Next
		dict[pHead] = clone
		pHead = pHead.Next
	}
	pClone = clonedDummy.Next
	for pClone != nil {
		pClone.Random = dict[pClone.Random]
		pClone = pClone.Next
	}
	return clonedDummy.Next
}
