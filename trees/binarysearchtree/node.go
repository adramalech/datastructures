package binarysearchtree

import "github.com/adramalech/datastructures/models"

type Node struct {
	left  *Node
	right *Node
	key   int
	value models.Person
}

func newNode(key int, value models.Person) *Node {
	return &Node{nil, nil, key, value}
}