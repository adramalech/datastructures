package singlelinkedlist

type SingleLinkedList struct {
	root *Node
}

func NewSingleLinkedList() SingleLinkedList {
	return SingleLinkedList{nil}
}

func (s *SingleLinkedList) Count() int {
	cnt := count(s.root)
	return cnt
}

func count(n *Node) int {
	if n == nil {
		return 0
	}

	return 1 + count(n.next)
}

func (s *SingleLinkedList) Insert(value int) {
	s.root = insert(s.root, value)
}

func insert(n *Node, value int) *Node {
	if n == nil {
		n = newNode(value)
		return n
	}

	n.next = insert(n.next, value)
	return n
}

func (s * SingleLinkedList) Find(value int) bool {
	found := find(s.root, value)
	return found
}

func find(n *Node, value int) bool {
	if n == nil {
		return false
	}

	if n.value == value {
		return true
	}

	return find(n.next, value)
}

func (s *SingleLinkedList) Remove(value int) {
	s.root = remove(s.root, value)
}

func remove(n *Node, value int) *Node {
	if n == nil {
		return nil
	}

	if n.value == value {
		n = n.next
		return n
	}

	n.next = remove(n.next, value)
	return n
}