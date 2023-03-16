package tree

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type NodeStack []*Node
type Visited []int

type Tree interface {
	Contains(int) bool
	Delete(int) *Node
	Get(int) *Node
	Insert(int) *Node
}

//func New(*Tree) *Tree {
//	return &Tree{
//		Root: nil,
//	}
//}

func CreateNode(value int) *Node {
	return &Node{
		Value: value,
		Left:  nil,
		Right: nil,
	}
}
