package queue

type Node struct {
	Data interface{}
	Next *Node
}

type Queue struct {
	First  *Node
	Last   *Node
	Length int
}

func New() *Queue {
	return &Queue{
		First:  nil,
		Last:   nil,
		Length: 0,
	}
}

func (q *Queue) Dequeue() *Node {
	if q.First == nil {
		return nil
	}
	node := q.First
	if q.Length == 1 {
		q.First = nil
		q.Last = nil
	} else {
		q.First = q.First.Next
		node.Next = nil
	}
	q.Length -= 1
	return node
}

func (q *Queue) Enqueue(data any) *Node {
	node := &Node{
		Data: data,
		Next: nil,
	}
	if q.First == nil {
		q.First = node
		q.Last = node
	} else {
		q.Last.Next = node
		q.Last = node
	}
	q.Length += 1
	return node
}

func (q *Queue) IsEmpty() bool {
	if q.Length == 0 {
		return true
	}
	return false
}

func (q *Queue) Reverse() *Node {
	if q.First == nil {
		return nil
	}
	current := q.First
	q.First = q.Last
	q.Last = current
	var prev *Node
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return q.First
}

func (q *Queue) ToArray() []int {
	if q.First == nil {
		return nil
	}
	current := q.Reverse()
	result := []int{}
	for current != nil {
		result = append(result, current.Data.(int))
		current = current.Next
	}
	return result
}
