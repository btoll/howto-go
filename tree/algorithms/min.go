package algorithms

import (
	"math"

	"github.com/btoll/howto-go/tree"
)

func r_min(node *tree.Node) float64 {
	if node == nil {
		return math.Inf(0)
	}
	return math.Min(float64(node.Value),
		math.Min(r_min(node.Left), r_min(node.Right)))
}

func i_min(node *tree.Node) int {
	if node == nil {
		return 0
	}
	min := 100
	stack := tree.NodeStack{node}
	var current *tree.Node
	for len(stack) > 0 {
		current, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if min > current.Value {
			min = current.Value
		}
		if current.Right != nil {
			stack = append(stack, current.Right)
		}
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
	}
	return min
}

func Min(nums tree.Visited) float64 {
	t := tree.CreateBT(nums)
	return r_min(t.Root)
	// return i_min(t.Root)
}
