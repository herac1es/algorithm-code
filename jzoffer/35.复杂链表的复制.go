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

// 1.先遍历链表，复制每一个节点，并将N'成为N的next
// 2.遍历，将复制节点的random更新为原值的next
// 3.分割两个链表，组装偶数节点为复制结果，并还原原来的链表
// 时间:O(n)
// 空间:O(1) 构建复制链表不算额外空间
func copyRandomListS2(head *Node) *Node {
	cloneNode(head)
	setRandom(head)
	return split(head)
}

// cloneNode 复制每一个节点，将复制出的N‘放在N的后面
func cloneNode(head *Node) {
	for head != nil {
		clone := &Node{
			Val:    head.Val,
			Next:   head.Next,
			Random: head.Random,
		}
		head.Next = clone
		head = clone.Next
	}
}

// setRandom 设置新节点的random指针
func setRandom(head *Node) {

	for head != nil {
		clone := head.Next
		if clone.Random != nil {
			clone.Random = clone.Random.Next
		}
		head = clone.Next
	}
}

// split 切分链表，并复原输入
// @return 复制后的链表
func split(head *Node) *Node {
	cloneDummy := &Node{}
	pClone := cloneDummy
	for head != nil {
		pClone.Next = head.Next
		pClone = pClone.Next
		next := head.Next.Next
		head.Next = next
		head = head.Next
	}
	return cloneDummy.Next
}
