package doublelinkedlist

type DoubleLinkedList struct {
	list *node
}

func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{nil}
}

func (d *DoubleLinkedList) InsertTail(value int) {
	d.list = insertNodeTail(d.list, value)
}

func (d *DoubleLinkedList) InsertHead(value int) {
	d.list = insertNodeHead(d.list, value)
}

func insertNodeTail(n *node, value int) *node {
	if n == nil {
		n = newNode(value)
		return n
	}

	n.next = insertNodeTail(n.next, value)
	return n
}

func insertNodeHead(n *node, value int) *node {
	if n == nil {
		n = newNode(value)
		return n
	}

	n.previous = insertNodeHead(n.previous, value)
	return n
}

func (d *DoubleLinkedList) Rewind() {
	d.list = rewind(d.list)
}

func rewind(n *node) *node {
	if n.previous == nil {
		return n
	}

	return rewind(n.previous)
}

func (d *DoubleLinkedList) FastForward() {
	d.list = fastForward(d.list)
}

func fastForward(n *node) *node {
	if n.next == nil {
		return n
	}

	return fastForward(n.next)
}