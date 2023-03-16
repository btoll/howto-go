// https://www.bigocheatsheet.com/
package linkedlist

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func New(value int) *LinkedList {
	node := &Node{
		Value: value,
		Next:  nil,
	}
	return &LinkedList{
		Head:   node,
		Tail:   node,
		Length: 1,
	}
}

func (l *LinkedList) Append(value int) bool {
	node := &Node{
		Value: value,
		Next:  nil,
	}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
		l.Tail = node
	}
	l.Length += 1
	return true
}

func (l *LinkedList) Get(index int) *Node {
	if index < 0 || index >= l.Length {
		return nil
	}
	if index == 0 {
		return l.Head
	}
	if index == l.Length-1 {
		return l.Tail
	}
	i := 0
	current := l.Head
	for i < index {
		current = current.Next
		i += 1
	}
	return current
}

func (l *LinkedList) Insert(index, value int) bool {
	if index < 0 || index > l.Length {
		return false
	}
	if index == 0 {
		return l.Prepend(value)
	}
	if index == l.Length {
		return l.Append(value)
	}
	node := &Node{
		Value: value,
		Next:  nil,
	}
	pre := l.Get(index - 1)
	node.Next = pre.Next
	pre.Next = node
	l.Length += 1
	return true
}

func (l *LinkedList) Pop() *Node {
	if l.Head == nil {
		return nil
	}
	current := l.Head
	if l.Head == l.Tail {
		l.Head = nil
		l.Tail = nil
	} else {
		var pre *Node
		for current.Next != nil {
			pre = current
			current = current.Next
		}
		l.Tail = pre
		l.Tail.Next = nil
	}
	l.Length -= 1
	return current
}

func (l *LinkedList) PopFirst() *Node {
	if l.Head == nil {
		return nil
	}
	current := l.Head
	if l.Head == l.Tail {
		l.Head = nil
		l.Tail = nil
	} else {
		l.Head = l.Head.Next
		current.Next = nil
	}
	l.Length -= 1
	return current
}

func (l *LinkedList) Prepend(value int) bool {
	node := &Node{
		Value: value,
		Next:  nil,
	}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		node.Next = l.Head
		l.Head = node
	}
	l.Length += 1
	return true
}

func (l *LinkedList) Print() {
	current := l.Head
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
}

func (l *LinkedList) Remove(index int) *Node {
	if index < 0 || index >= l.Length {
		return nil
	}
	if index == 0 {
		return l.PopFirst()
	}
	if index == l.Length-1 {
		return l.Pop()
	}
	prev := l.Get(index - 1)
	target := prev.Next
	prev.Next = target.Next
	target.Next = nil
	l.Length -= 1
	return target
}

func (l *LinkedList) Reverse() {
	current := l.Head
	l.Head = l.Tail
	l.Tail = current

	var prev *Node
	for current != nil {
		// Save the next node in a local var b/c we're
		// about to unlink it and without this reference
		// we're fucked.
		next := current.Next
		current.Next = prev
		// Advance both nodes.
		prev = current
		current = next
	}
}

func (l *LinkedList) Set(index, value int) bool {
	current := l.Get(index)
	if current != nil {
		current.Value = value
		return true
	}
	return false
}
