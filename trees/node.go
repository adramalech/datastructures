package trees

type Node struct {
	left  *Node
	right *Node
	key int
}

func NewNode(key int) *Node {
	return &Node{nil, nil, key}
}
