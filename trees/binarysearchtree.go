package trees

type Tree interface {
	Insert(key int)
	Exists(key int) bool
	Max() int
	Min() int
	Count() int
	Delete(int) *Node
}

func (n *Node) isLeftNil() bool {
	return n.left == nil
}

func (n *Node) isRightNil() bool {
	return n.right == nil
}

func (n *Node) Insert(key int) {
	if key == n.key {
		// we skip
		return
	}

	if key > n.key {
		if n.isRightNil() {
			n.right = NewNode(key)
		} else {
			n.right.Insert(key)
		}
	} else {
		if n.isLeftNil() {
			n.left = NewNode(key)
		} else {
			n.left.Insert(key)
		}
	}
}

func (n *Node) Min() int {
	// if left is nil and right is greater than current then return current key.
	if n.isLeftNil() {
		return n.key
	}

	return n.left.Min()
}

func (n *Node) Count() int {
	if n == nil {
		return 0
	}

	// if at a leaf return the count of the leaf.
	if n.isLeftNil() && n.isRightNil() {
		return 1
	}

	// if current node children both exist
	if !n.isLeftNil() && !n.isRightNil() {
		return 1 + n.left.Count() + n.right.Count()
	}

	// if right is nil go left
	if n.isRightNil() {
		return 1 + n.left.Count()
	}

	// if left is nil go right
	if n.isLeftNil() {
		return 1 + n.right.Count()
	}

	return 0
}

func (n *Node) Max() int {
	// if right is nil return current because left will be less.
	if n.isRightNil() {
		return n.key
	}

	return n.right.Max()
}

func (n *Node) Exists(key int) bool {
	if n == nil {
		return false
	}

	// if we found the key return it.
	if n.key == key {
		return true
	}

	// we are at a leaf and have not found the key return nothing.
	if n.isLeftNil() && n.isRightNil() {
		return false
	}

	// if current key is greater than current key search right.
	if key > n.key {
		return n.right.Exists(key)
	}

	// else current key is less than current key search left.
	return n.left.Exists(key)
}

func (n *Node) Delete(key int)  *Node {
	// key isn't found.
	if n == nil {
		return n
	}

	// key is less than current node travel left.
	if key < n.key {
		n.left = n.left.Delete(key)
	}

	// key is greater than current node travel right
	if key > n.key {
		n.right = n.right.Delete(key)
	}

	if key == n.key {
		if !n.isLeftNil() && !n.isRightNil() {
			//set N key to minimum of right.
			n.key = n.right.Min()

			//delete the node of the right side.
			n.right = n.right.Delete(n.key)
		}

		if n.isRightNil() {
			return n.left
		} else {
			return n.right
		}
	}

	return n
}
