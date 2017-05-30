package main

func (t *Tree) FindRecursive(n *Node, value int) *Node {
	if n == nil {
		return nil
	}

	if n.Value == value {
		return n
	}

	if value < n.Value {
		return t.FindRecursive(n.Left, value)
	} else {
		return t.FindRecursive(n.Right, value)
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
