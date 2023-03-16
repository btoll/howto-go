package tree

type BST struct {
	Root *Node
}

func NewBST() *BST {
	return &BST{
		Root: nil,
	}
}

func CreateBST(nums Visited) *BST {
	t := NewBST()
	for _, n := range nums {
		t.Insert(n)
	}
	return t
}

func (b *BST) Contains(value int) bool {
	node := b.Root
	for node != nil {
		if value == node.Value {
			return true
		}
		if value < node.Value {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return false
}

//	              10
//	            /    \
//	           /      \
//	          7        25
//	         / \      / \
//	        /   \    /   \
//	       4     8  21    32
//		    \			 /  \
//			 \		    /    \
//			  5		  30	  40
func (b *BST) i_delete(value int) *Node {
	if b.Root == nil {
		return nil
	}
	var prev *Node
	current := b.Root
	for current != nil && value != current.Value {
		//		if value == current.Value {
		//			break
		//		}
		prev = current
		if value < current.Value {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	// Node isn't in the tree
	if current == nil {
		return nil
	}
	//	fmt.Println("prev", prev)
	//	fmt.Println("current", current)
	if current.Left == nil || current.Right == nil {
		var n *Node = nil
		if current.Left != nil {
			n = current.Left
		} else {
			n = current.Right
		}
		// This is akin to removing a node in a linked list.
		if current == prev.Left {
			prev.Left = n
		} else {
			prev.Right = n
		}
	} else {
		minNode := b.MinValue(current.Right)
		b.Delete(minNode.Value)
		current.Value = minNode.Value
	}
	return nil
}

func (b *BST) r_delete(node *Node, value int) *Node {
	if node == nil {
		return nil
	}
	if value < node.Value {
		node.Left = b.r_delete(node.Left, value)
	}
	if value > node.Value {
		node.Right = b.r_delete(node.Right, value)
	}
	if value == node.Value {
		if node.Left == nil && node.Right == nil {
			return nil
		}
		if node.Left != nil && node.Right == nil {
			return node.Left
		}
		if node.Left == nil && node.Right != nil {
			return node.Right
		}
		if node.Left != nil && node.Right != nil {
			minNode := b.MinValue(node)
			b.r_delete(node, minNode.Value)
			node.Value = minNode.Value
			return node
		}
	}
	return node
}

func (b *BST) Delete(value int) *Node {
	//	return b.r_delete(b.Root, value)
	return b.i_delete(value)
}

func (b *BST) Get(value int) *Node {
	if b.Root == nil {
		return nil
	}
	current := b.Root
	for current != nil {
		if value == current.Value {
			return current
		}
		if value < current.Value {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	return nil
}

func (b *BST) Insert(value int) *Node {
	node := CreateNode(value)
	if b.Root == nil {
		b.Root = node
		return node
	}
	current := b.Root
	for {
		if value == current.Value {
			return nil
		}
		if value < current.Value {
			if current.Left == nil {
				current.Left = node
				return node
			}
			current = current.Left
		} else {
			if current.Right == nil {
				current.Right = node
				return node
			}
			current = current.Right
		}
	}
	return nil
}

func (b *BST) MinValue(node *Node) *Node {
	if node == nil {
		return nil
	}
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}
