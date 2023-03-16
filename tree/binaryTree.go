package tree

import "github.com/btoll/howto-go/queue"

type BT struct {
	Root *Node
}

func NewBT() *BT {
	return &BT{
		Root: nil,
	}
}

func CreateBT(nums Visited) *BT {
	t := NewBT()
	for _, n := range nums {
		t.Insert(n)
	}
	return t
}

//	              _1_
//	            /     \
//	           /       \
//	          2         3
//	         / \       / \
//	        /   \     /   \
//	       4     5   6     7
//	     /  \   /
//		/   |  |
//	   8    9  10

func (b *BT) Insert(value int) *Node {
	node := CreateNode(value)
	if b.Root == nil {
		b.Root = node
		return node
	}
	q := queue.New()
	q.Enqueue(b.Root)
	var current *queue.Node
	var cnode *Node
	for !q.IsEmpty() {
		current = q.Dequeue()
		cnode = current.Data.(*Node)
		if cnode.Left == nil {
			cnode.Left = node
			return node
		} else {
			q.Enqueue(cnode.Left)
		}
		if cnode.Right == nil {
			cnode.Right = node
			return node
		} else {
			q.Enqueue(cnode.Right)
		}
	}
	return nil
}
