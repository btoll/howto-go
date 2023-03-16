package stack

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type Stack struct {
	Top    *Node
	Height int
}

func New(value int) *Stack {
	return &Stack{
		Top: &Node{
			Value: value,
			Next:  nil,
		},
		Height: 0,
	}
}

func (s *Stack) Pop() *Node {
	if s.Pop == nil {
		return nil
	}
	node := s.Top
	s.Top = s.Top.Next
	node.Next = nil
	s.Height -= 1
	return node
}

func (s *Stack) Print() {
	current := s.Top
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
}

func (s *Stack) Push(value int) {
	node := &Node{
		Value: value,
		Next:  nil,
	}
	if s.Top == nil {
		s.Top = node
	} else {
		node.Next = s.Top
		s.Top = node
	}
	s.Height += 1
}
