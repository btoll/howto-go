package algorithms

import (
	"github.com/btoll/howto-go/tree"
)

/*
//	       10
//	     /    \
//	    /      \
//	   7        25
//	  / \      / \
//	 /   \    /   \
//	4     8  21    32
*/

func r_maxpath(node *tree.Node, pathSum int) int {
	if node == nil {
		return pathSum
	}
	pathSum += node.Value
	left := r_maxpath(node.Left, pathSum)
	right := r_maxpath(node.Right, pathSum)
	if right > left {
		return right
	}
	return left
}

func MaxPath(nums []int) int {
	t := tree.CreateBT(nums)
	return r_maxpath(t.Root, 0)
}
