package tree

import (
	"github.com/btoll/howto-go/queue"
)

// Bfs(*Node) Visited
// Inorder(*Node) Visited
// Levelorder(node *Node) []Visited
// Postorder(*Node) Visited
// Preorder(*Node) Visited

//	       10
//	     /    \
//	    /      \
//	   7        25
//	  / \      / \
//	 /   \    /   \
//	4     8  21    32
//
// bfs       [10, 7, 25, 4, 8, 21, 32]
// preorder  [10, 7, 4, 8, 25, 21, 32]
// inorder   [4, 7, 8, 10, 21, 25, 32]
// postorder [4, 8, 7, 21, 32, 25, 10]

func Bfs(node *Node) Visited {
	if node == nil {
		return nil
	}
	q := queue.New()
	q.Enqueue(node)
	visited := Visited{}
	var current *Node
	for !q.IsEmpty() {
		current = q.Dequeue().Data.(*Node)
		visited = append(visited, current.Value)
		if current.Left != nil {
			q.Enqueue(current.Left)
		}
		if current.Right != nil {
			q.Enqueue(current.Right)
		}
	}
	return visited
}

func Inorder(node *Node) Visited {
	if node == nil {
		return nil
	}
	stack := NodeStack{}
	visited := Visited{}
	current := node
	for true {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		if len(stack) == 0 {
			break
		}
		current, stack = stack[len(stack)-1], stack[:len(stack)-1]
		visited = append(visited, current.Value)
		current = current.Right
	}
	return visited
}

func Levelorder(node *Node) []Visited {
	if node == nil {
		return nil
	}
	q := queue.New()
	q.Enqueue(node)
	levels := []Visited{}
	var current *Node
	for !q.IsEmpty() {
		size := q.Length
		level := Visited{}
		for i := 0; i < size; i += 1 {
			current = q.Dequeue().Data.(*Node)
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

func Postorder(node *Node) Visited {
	stack := NodeStack{node}
	q := queue.New()
	var current *Node
	for len(stack) > 0 {
		current, stack = stack[len(stack)-1], stack[:len(stack)-1]
		q.Enqueue(current.Value)
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
		if current.Right != nil {
			stack = append(stack, current.Right)
		}
	}
	return q.ToArray()
}

func Preorder(node *Node) Visited {
	if node == nil {
		return nil
	}
	stack := NodeStack{node}
	visited := Visited{}
	var current *Node
	for len(stack) > 0 {
		current, stack = stack[len(stack)-1], stack[:len(stack)-1]
		visited = append(visited, current.Value)
		if current.Right != nil {
			stack = append(stack, current.Right)
		}
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
	}
	return visited
}
