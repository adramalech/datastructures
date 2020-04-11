package binarysearchtree

import "github.com/adramalech/datastructures/models"

type BinarySeachTree struct {
	root *Node
}

func NewBinarySearchTree() *BinarySeachTree {
	return &BinarySeachTree{nil}
}

func (bst *BinarySeachTree) Insert(key int, value models.Person) {
	if bst.root == nil {
		bst.root = newNode(key, value)
	}

	if key == bst.root.key {
		// we skip
		return
	}

	bst.root = insertNode(bst.root, key, value)
}

func (bst *BinarySeachTree) Count() int {
	total := countNode(bst.root)
	return total
}

func (bst *BinarySeachTree) KeyExists(key int) bool {
	exists := nodeExists(bst.root, key)
	return exists
}

func (bst *BinarySeachTree) Delete(key int) {
	bst.root = deleteNode(bst.root, key)
}

// binary search tree helpers

func countNode(n *Node) int {
	if n == nil {
		return 0
	}

	// leaf
	if n.left == nil && n.right == nil {
		return 1
	}

	return 1 + countNode(n.left) + countNode(n.right)
}

func deleteNode(n *Node, key int) *Node {
	if n == nil {
		return nil
	}

	if key > n.key {
		n.right = deleteNode(n.right, key)
		return n
	}

	if key < n.key {
		n.left = deleteNode(n.left, key)
		return n
	}

	// this is a leaf delete it.
	if n.left == nil && n.right == nil {
		n = nil
		return n
	}

	// left is nil move right up
	if n.left == nil {
		n = n.right
		return n
	}

	// right is nil move left up
	if n.right == nil {
		n = n.left
		return n
	}

	// both children are not nil
	minKeyNode := minNode(n.right)
	n.key = minKeyNode.key
	n.value = minKeyNode.value
	n.right = deleteNode(n.right, key)

	return n
}

func insertNode(node *Node, key int, value models.Person) *Node {
	if node == nil {
		node = newNode(key, value)
		return node
	}

	if key > node.key {
		node.right = insertNode(node.right, key, value)
		return node
	}

	if key < node.key {
		node.left = insertNode(node.left, key, value)
		return node
	}

	return node
}

func nodeExists(n *Node, key int) bool {
	if n == nil {
		return false
	}

	if key == n.key {
		return true
	}

	if key > n.key {
		return nodeExists(n.right, key)
	}

	return nodeExists(n.left, key)
}

func minNode(n *Node) *Node {
	if n.left == nil {
		return n
	}

	return minNode(n.left)
}


// end binary search tree helpers
