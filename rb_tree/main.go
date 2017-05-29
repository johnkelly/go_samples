package main

import (
	"fmt"

	"github.com/fatih/color"
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

type DisplayNode struct {
	Thing *Node
	Level int
}

// Works for very small single digit trees
func (t *Tree) Display() {
	fmt.Printf("Max Depth: %d\n\n", t.Root.maxDepth())

	maxLevel := t.Root.maxDepth() - 1
	lastLevel := 0

	var queue []*DisplayNode
	level := 0
	queue = append(queue, &DisplayNode{Thing: t.Root, Level: level})

	for len(queue) != 0 && lastLevel <= maxLevel {
		element := queue[0]
		queue = queue[1:]

		node := element.Thing
		level := element.Level

		if node.Left != nil {
			newNode := &DisplayNode{
				Thing: node.Left,
				Level: level + 1,
			}
			queue = append(queue, newNode)
		} else {
			newNode := &DisplayNode{
				Thing: &Node{Value: -1, Red: false},
				Level: level + 1,
			}
			queue = append(queue, newNode)
		}

		if node.Right != nil {
			newNode := &DisplayNode{
				Thing: node.Right,
				Level: level + 1,
			}
			queue = append(queue, newNode)
		} else {
			newNode := &DisplayNode{
				Thing: &Node{Value: -1, Red: false},
				Level: level + 1,
			}
			queue = append(queue, newNode)
		}

		if level != lastLevel {
			fmt.Println("")
			lastLevel = level
		}
		node.display(level, maxLevel)
	}
}

func (n *Node) display(level, maxLevel int) {
	red := color.New(color.FgRed)
	black := color.New(color.FgBlack)
	whiteRed := red.Add(color.BgWhite)
	whiteBlack := black.Add(color.BgWhite)

	msg := ""
	numTabs := (maxLevel - level) + 1

	for i := 0; i < numTabs; i++ {
		msg += "\t"
	}
	fmt.Print(msg)
	if n.Red {
		whiteRed.Printf("%d", n.Value)
	} else {
		whiteBlack.Printf("%d", n.Value)
	}
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
