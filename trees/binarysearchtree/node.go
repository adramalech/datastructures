package binarysearchtree

import "github.com/adramalech/datastructures/models"

type node struct {
	left  *node
	right *node
	key   int
	value models.Person
}

func newNode(key int, value models.Person) *node {
	return &node{nil, nil, key, value}
}