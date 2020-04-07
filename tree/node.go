package trees

type Node struct {
	left  *Node
	right *Node
	value int
}

func NewNode(value int) *Node {
	return &Node{nil, nil, value}
}
