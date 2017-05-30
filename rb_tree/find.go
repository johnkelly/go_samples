package main

func (t *Tree) FindRecursive(value int) *Node {
	return t.FindRecursiveHelper(t.Root, value)
}

func (t *Tree) FindRecursiveHelper(n *Node, value int) *Node {
	if n == nil {
		return nil
	}

	if n.Value == value {
		return n
	}

	if value < n.Value {
		return t.FindRecursiveHelper(n.Left, value)
	} else {
		return t.FindRecursiveHelper(n.Right, value)
	}
	return nil
}

func (t *Tree) FindIterative(value int) *Node {
	if t.Root == nil {
		return nil
	}
	node := t.Root

	for node != nil {
		if node.Value == value {
			return node
		}

		if value < node.Value {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return nil
}
