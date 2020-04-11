package singlelinkedlist

type SingleLinkedList struct {
	root *node
}

func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{nil}
}

func (s *SingleLinkedList) Count() int {
	cnt := count(s.root)
	return cnt
}

func count(n *node) int {
	if n == nil {
		return 0
	}

	return 1 + count(n.next)
}

func (s *SingleLinkedList) Insert(value int) {
	s.root = insertNode(s.root, value)
}

func insertNode(n *node, value int) *node {
	if n == nil {
		n = newNode(value)
		return n
	}

	n.next = insertNode(n.next, value)
	return n
}

func (s *SingleLinkedList) Find(value int) bool {
	found := findNode(s.root, value)
	return found
}

func findNode(n *node, value int) bool {
	if n == nil {
		return false
	}

	if n.value == value {
		return true
	}

	return findNode(n.next, value)
}

func (s *SingleLinkedList) Remove(value int) {
	s.root = removeNode(s.root, value)
}

func removeNode(n *node, value int) *node {
	if n == nil {
		return nil
	}

	if n.value == value {
		n = n.next
		return n
	}

	n.next = removeNode(n.next, value)
	return n
}