package singlelinkedlist

type Node struct {
	next *Node
	value int
}

func newNode(value int) *Node {
	return &Node{nil, value}
}