package doublelinkedlist

type node struct {
	previous *node
	next *node
	value int
}

func newNode(value int) *node {
	return &node{nil, nil, value}
}