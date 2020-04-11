package singlelinkedlist

type node struct {
	next *node
	value int
}

func newNode(value int) *node {
	return &node{nil, value}
}