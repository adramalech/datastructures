package trees

type Tree interface {
	Insert(value int)
	Find(value int) bool
	Remove(value int)
	Max() *Node
	Min() *Node
}

func (n *Node) GetValue() int {
	return n.value
}

func (n *Node) isLeftNil() bool {
	return n.left == nil
}

func (n *Node) isRightNil() bool {
	return n.right == nil
}

func (n *Node) Insert(value int) {
	if value == n.value {
		// we skip
		return
	}

	// great value add to right.
	if value > n.value {
		if n.isRightNil() {
			n.right = NewNode(value)
		} else {
			n.right.Insert(value)
		}
	} else {
		if n.isLeftNil() {
			n.left = NewNode(value)
		} else {
			n.left.Insert(value)
		}
	}
}

func (n *Node) Min() *Node {
	// at a dead end return value
	if n.isLeftNil() && n.isRightNil() {
		return n
	}

	// if left is nil and right is greater than current then return current value.
	if n.isLeftNil() {
		return n
	}

	return n.left.Min()
}

func (n *Node) removeMinNode() *Node {
	return nil
}

func (n *Node) Max() *Node {
	// at a dead end return value
	if n.isLeftNil() && n.isRightNil() {
		return n
	}

	// if right is nil return current because left will be less.
	if n.isRightNil() {
		return n
	}

	return n.right.Max()
}

func (n *Node) Find(value int) bool {
	// if we found the value return it.
	if n.value == value {
		return true
	}

	// we are at a leaf and have not found the value return nothing.
	if n.isLeftNil() && n.isRightNil() {
		return false
	}

	// if current value is greater than current value search right.
	if value > n.value {
		return n.right.Find(value)
	}

	// else current value is less than current value search left.
	return n.left.Find(value)
}

func (n *Node) Remove(value int) {
	// the value is found.
	if n.value == value {
		// if leaf remove
		if n.isLeftNil() && n.isRightNil() {
			n = nil
			return
		}

		// move right child up if left is empty
		if n.isLeftNil() {
			n = n.right
			return
		}

		// move left child up if right is empty
		if n.isRightNil() {
			n = n.left
			return
		}

		// neither is empty move up child with minimum greatest value.
		if !n.isLeftNil() && !n.isRightNil() {
			var minNode *Node = n.right.Min()
			var rightChild *Node = n.right
			var leftChild *Node = n.left

			n = minNode
			n.right = rightChild
			n.left = leftChild
		}

		return
	}

	// if we are at leaf return (dead end = "nothing found")
	if n.isLeftNil() && n.isRightNil() {
		return
	}

	// if value is greater than current value go right
	// else value is left go left.
	if value > n.value {
		n.right.Remove(value)
	} else {
		n.left.Remove(value)
	}
}
