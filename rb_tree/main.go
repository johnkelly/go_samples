package main

import (
	"fmt"
	"os"
)

type Tree struct {
	Root *Node
}

type Node struct {
	Parent, Left, Right *Node
	Red                 bool
	Value               int
}

func main() {
	input := []int{3, 1, 5, 7, 6, 8, 9, 10, 11, 20}

	tree := &Tree{}
	tree.InsertAll(input)
	tree.Display()
	fmt.Printf("Find Results for 4: %v\n", tree.FindIterative(4))
	fmt.Printf("Find Results for 10: %v\n", tree.FindIterative(10))

	fmt.Printf("Find Results for 4: %v\n", tree.FindRecursive(4))
	fmt.Printf("Find Results for 10: %v\n", tree.FindRecursive(10))

	f, err := os.OpenFile("tree.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 777)
	if err != nil {
		panic("Can't open file to write tree")
	}
	defer f.Close()

	tree.Serialize(tree.Root, f)

	f, err = os.OpenFile("tree.txt", os.O_RDONLY, 777)
	if err != nil {
		panic("Can't open file to write tree")
	}
	defer f.Close()

	var root *Node
	root = DeSerialize(root, f)
	tree = &Tree{Root: root}
	tree.Display()
}

func (t *Tree) Serialize(n *Node, f *os.File) {
	if n == nil {
		fmt.Fprintf(f, "%d|%t,", -1, false)
		return
	}
	fmt.Fprintf(f, "%d|%t,", n.Value, n.Red)

	t.Serialize(n.Left, f)
	t.Serialize(n.Right, f)
}

func DeSerialize(n *Node, f *os.File) *Node {
	var value int
	var red bool

	_, err := fmt.Fscanf(f, "%d|%t,", &value, &red)
	if err != nil {
		return nil
	}

	if value == -1 {
		return nil
	}

	n = &Node{Value: value, Red: red}

	n.Left = DeSerialize(n.Left, f)
	n.Right = DeSerialize(n.Right, f)

	return n
}

func (t *Tree) InsertAll(values []int) {
	for _, number := range values {
		t.Insert(number)
	}
}

func (t *Tree) Insert(value int) {
	x := t.Root
	var y *Node

	// Advance the x pointer to where in the
	// graph value belongs (iteration)
	for x != nil {
		y = x
		if value < x.Value {
			x = x.Left
		} else {
			x = x.Right
		}
	}

	// Create the new node - new node's are always red
	newNode := &Node{Parent: y, Value: value, Red: true}

	// Insert the new Node into the graph
	if y == nil {
		// The tree is empty
		newNode.Red = false
		t.Root = newNode
		return
	} else if newNode.Value < y.Value {
		// The new value goes on the left
		y.Left = newNode
	} else {
		// The new value goes on the right
		y.Right = newNode
	}
	// Rebalance the tree
	t.Balance(newNode)
}

func (t *Tree) Balance(problemNode *Node) {
	var aunt *Node
	// Advance up the tree to check for violations
	for problemNode.Parent != nil && problemNode.Parent.Red {
		if problemNode.Parent == problemNode.Parent.Parent.Left {
			aunt = problemNode.Parent.Parent.Right
			if aunt != nil && aunt.Red {
				problemNode = redAunt(problemNode, aunt)
			} else {
				t.blackAunt(problemNode, true)
			}
		} else {
			aunt = problemNode.Parent.Parent.Left
			if aunt != nil && aunt.Red {
				problemNode = redAunt(problemNode, aunt)
			} else {
				t.blackAunt(problemNode, false)
			}
		}
	}
	t.Root.Red = false
}

// Black Aunt
func (t *Tree) blackAunt(x *Node, left bool) {
	// if the work is being down on the left side
	// of the tree else we are on the right side
	if left {
		// parent
		//  / \
		//     x
		// In this case left rotate
		// bringing grand parent down
		// and promoting parent up the tree
		if x == x.Parent.Right {
			x = x.Parent
			t.leftRotate(x)
		}
		// Same idea as color flip red grand parent
		// black parents and aunts
		// we know that the aunt is already black so
		// we only need to color flip the parent &
		// grandparent
		x.Parent.Red = false
		x.Parent.Parent.Red = true

		// Right Rotate the grandparent as we are
		// on the left side of the tree
		t.rightRotate(x.Parent.Parent)
	} else {
		// parent
		//  / \
		// x
		// In this case right rotate
		// bringing grand parent down
		// and promoting parent up the tree
		if x == x.Parent.Left {
			x = x.Parent
			t.rightRotate(x)
		}
		// Same idea as color flip red grand parent
		// black parents and aunts
		// we know that the aunt is already black so
		// we only need to color flip the parent &
		// grandparent
		x.Parent.Red = false
		x.Parent.Parent.Red = true

		// Left Rotate the grandparent as we are
		// on the right side of the tree
		t.leftRotate(x.Parent.Parent)
	}
}

// Red Aunt
// Color Flip and then move up the tree
// to check for further violations
func redAunt(x, aunt *Node) *Node {
	colorFlip(x, aunt)
	return x.Parent.Parent
}

// Color Flip
//    red
//    /  \
// black black
// set the Grandparent as red
// set the aunts & parents to black
func colorFlip(x, aunt *Node) {
	x.Parent.Red = false
	aunt.Red = false
	x.Parent.Parent.Red = true
}

// Rotates the subtree counter clockwise
//     1
//       \
//         2
//          \
//           3
func (t *Tree) leftRotate(x *Node) {
	// x is 1

	// assign a pointer to the 2
	y := x.Right

	// set the right of 1 to null
	// 1
	//
	// 2
	//  \
	//   3
	x.Right = y.Left

	if y.Left != nil {
		// if 2 has a child on left
		// make it's parent 1
		y.Left.Parent = x
	}
	// set the parent of 1 to be the parent of 2
	// 2
	//  \
	//   3
	//
	// 1
	y.Parent = x.Parent
	if x.Parent == nil {
		// connect the 2\3 subtree to the rest
		// of the tree as the root
		t.Root = y
	} else if x == x.Parent.Left {
		// connect the 2\3 subtree to the rest
		// of the tree on the left side
		x.Parent.Left = y
	} else {
		// connect the 2\3 subtree to the rest
		// of the tree on the right side
		x.Parent.Right = y
	}
	// reconnect the 1 subtree to the 2/3
	// subtree on the left side
	//      2
	//     / \
	//    1   3
	y.Left = x
	x.Parent = y
}

// Rotates the subtree clockwise
//						3
//					/
//				2
//			/
//		1
func (t *Tree) rightRotate(x *Node) {
	// x is 3

	// assign a pointer to the 2
	y := x.Left

	// set the left of 3 to null
	// 3
	//
	//     2
	//   /
	//  1
	x.Left = y.Right

	if y.Right != nil {
		// if 2 has a child on right
		// make it's parent 3
		y.Right.Parent = x
	}
	// set the parent of 3 to be the parent of 2
	//     2
	//   /
	//  1
	//
	// 3
	y.Parent = x.Parent
	if x.Parent == nil {
		// connect the 2/1 subtree to the rest
		// of the tree as the root
		t.Root = y
	} else if x == x.Parent.Left {
		// connect the 2\1 subtree to the rest
		// of the tree on the left side
		x.Parent.Left = y
	} else {
		// connect the 2\1 subtree to the rest
		// of the tree on the right side
		x.Parent.Right = y
	}
	// reconnect the 3 subtree to the 2/1
	// subtree on the right side
	//      2
	//     / \
	//    1   3
	y.Right = x
	x.Parent = y
}

func (n *Node) maxDepth() int {
	if n == nil {
		return 0
	}
	leftDepth := n.Left.maxDepth()
	rightDepth := n.Right.maxDepth()

	if leftDepth > rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}
