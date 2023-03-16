package algorithms

import (
	"github.com/btoll/howto-go/queue"
	"github.com/btoll/howto-go/tree"
)

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

// [[1] [2 3] [4 5 6 7] [8 9 10]]

func levelOrder(node *tree.Node, levels []tree.Visited, level int) []tree.Visited {
	if node == nil {
		return levels
	}
	if len(levels) == level {
		levels = append(levels, tree.Visited{})
	}
	levels[level] = append(levels[level], node.Value)
	levels = levelOrder(node.Left, levels, level+1)
	levels = levelOrder(node.Right, levels, level+1)
	return levels
}

func LevelOrder(node *tree.Node) []tree.Visited {
	if node == nil {
		return nil
	}
	q := queue.New()
	q.Enqueue(node)
	levels := []tree.Visited{}
	for !q.IsEmpty() {
		size := q.Length
		level := tree.Visited{}
		for i := 0; i < size; i += 1 {
			current := q.Dequeue().Data.(*tree.Node)
			level = append(level, current.Value)
			if current.Left != nil {
				q.Enqueue(current.Left)
			}
			if current.Right != nil {
				q.Enqueue(current.Right)
			}
		}
		levels = append(levels, level)
	}
	return levels
}

func ListOfDepths(nums tree.Visited) []tree.Visited {
	t := tree.CreateBT(nums)
	//	return levelOrder(t.Root, []tree.Visited{}, 0)
	return LevelOrder(t.Root)
}
