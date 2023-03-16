package algorithms

import (
	"github.com/btoll/howto-go/tree"
)

func r_sum(node *tree.Node) int {
	if node == nil {
		return 0
	}
	return node.Value + r_sum(node.Left) + r_sum(node.Right)
}

func i_sum(node *tree.Node) int {
	if node == nil {
		return 0
	}
	total := 0
	stack := tree.NodeStack{node}
	var current *tree.Node
	for len(stack) > 0 {
		current, stack = stack[len(stack)-1], stack[:len(stack)-1]
		total += current.Value
		if current.Right != nil {
			stack = append(stack, current.Right)
		}
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
	}
	return total
}

func Sum(nums tree.Visited) int {
	t := tree.CreateBT(nums)
	return r_sum(t.Root)
	// return i_sum(t.Root)
}
