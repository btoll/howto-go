package algorithms

import (
	"github.com/btoll/howto-go/tree"
)

// Solution #1:
//		- do an inorder traversal
//		- check if slice is sorted
//		- only works if there are no duplicates
//
// If there are duplicates, this fails because there is no way to distinguish
// between the following:
//
//		42			42
//	   /			  \
//	 42			       42
//
//	(only the one on the left is valid)

//func ValidateBST(nums []int) bool {
//	t := tree.CreateBST(nums)
//	inorder := tree.Inorder(t.Root)
//	for i := 1; i < len(inorder); i++ {
//		if inorder[i] <= inorder[i-1] {
//			return false
//		}
//	}
//	return true
//}

// Solution #2:
//		- do an inorder traversal, but don't append values to a slice
//		- there's no need to, just compare the current value against the last one
//		- works if there are duplicates (testing <= against the last value)

//func ValidateBST(nums []int) bool {
//	t := tree.CreateBST(nums)
//	stack := tree.NodeStack{}
//	current := t.Root
//	// We want nil to be original value to compare against
//	// because any number could be in the tree.
//	var lastValueSeen *int
//	for {
//		for current != nil {
//			stack = append(stack, current)
//			current = current.Left
//		}
//		if len(stack) == 0 {
//			break
//		}
//		// Slice pop() idiom.
//		current, stack = stack[len(stack)-1], stack[:len(stack)-1]
//		if lastValueSeen != nil && current.Value <= *lastValueSeen {
//			return false
//		}
//		lastValueSeen = &current.Value
//		current = current.Right
//	}
//	return true
//}

// Solution #3:
//		- do an inorder traversal, but don't append values to a slice or append to a stack
//		- there's no need to, just ensure that every node to the left is less than and
//		  every node to the right is greater than the root node of the subtree
//		- works if there are duplicates

//		              10
//		            /    \
//		           /      \
//		          7        25
//		         / \      / \
//		        /   \    /   \
//		       4     8  21    32
//	 		    \			/  \
//			     \		   /    \
//				  5		 30		 40
var min *int
var max *int

func _validate(node *tree.Node, min, max *int) bool {
	if node == nil {
		return true
	}
	if (min != nil && node.Value <= *min) || (max != nil && node.Value > *max) {
		return false
	}
	if !_validate(node.Left, min, &node.Value) || !_validate(node.Right, &node.Value, max) {
		return false
	}
	return true
}

func ValidateBST(nums []int) bool {
	t := tree.CreateBST(nums)
	return _validate(t.Root, min, max)
}
