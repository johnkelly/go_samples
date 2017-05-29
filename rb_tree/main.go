package main

import (
	"fmt"
	"math"

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
	input := []int{3, 1, 5, 7, 6, 8, 9, 10}

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

	for x != nil {
		y = x
		if value < x.Value {
			x = x.Left
		} else {
			x = x.Right
		}
	}

	problemNode := &Node{Parent: y, Value: value, Red: true}

	if y == nil {
		problemNode.Red = false
		t.Root = problemNode
		return
	} else if problemNode.Value < y.Value {
		y.Left = problemNode
	} else {
		y.Right = problemNode
	}
	t.Balance(problemNode)
}

func (t *Tree) Balance(problemNode *Node) {
	var aunt *Node
	for problemNode.Parent != nil && problemNode.Parent.Red {
		if problemNode.Parent == problemNode.Parent.Parent.Left {
			aunt = problemNode.Parent.Parent.Right
			if aunt != nil && aunt.Red {
				// Red Aunt
				// Color Flip
				problemNode.Parent.Red = false
				aunt.Red = false
				problemNode.Parent.Parent.Red = true
				problemNode = problemNode.Parent.Parent
			} else {
				// Black Aunt
				// Rotate
				if problemNode == problemNode.Parent.Right {
					problemNode = problemNode.Parent
					t.leftRotate(problemNode)
				}
				problemNode.Parent.Red = false
				problemNode.Parent.Parent.Red = true
				t.rightRotate(problemNode.Parent.Parent)
			}
		} else {
			aunt = problemNode.Parent.Parent.Left
			if aunt != nil && aunt.Red {
				// Red Aunt
				//Color Flip
				problemNode.Parent.Red = false
				aunt.Red = false
				problemNode.Parent.Parent.Red = true
				problemNode = problemNode.Parent.Parent
			} else {
				// Black Aunt
				// Rotate
				if problemNode == problemNode.Parent.Left {
					problemNode = problemNode.Parent
					t.rightRotate(problemNode)
				}
				problemNode.Parent.Red = false
				problemNode.Parent.Parent.Red = true
				t.leftRotate(problemNode.Parent.Parent)
			}
		}
	}
	t.Root.Red = false
}

// Rotates the subtree counter clockwise
func (t *Tree) leftRotate(x *Node) {
	y := x.Right
	x.Right = y.Left

	if y.Left != nil {
		y.Left.Parent = x
	}
	y.Parent = x.Parent
	if x.Parent == nil {
		t.Root = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}
	y.Left = x
	x.Parent = y
}

// Rotates the subtree clockwise
func (t *Tree) rightRotate(x *Node) {
	y := x.Left
	x.Left = y.Right

	if y.Right != nil {
		y.Right.Parent = x
	}
	y.Parent = x.Parent
	if x.Parent == nil {
		t.Root = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}
	y.Right = x
	x.Parent = y
}

type DisplayNode struct {
	Thing *Node
	Level int
}

// Build the queue outside of the display
// then display all of the queue
func (t *Tree) Display() {
	fmt.Printf("Max Depth: %d\n\n", t.Root.maxDepth())

	maxLevel := t.Root.maxDepth() - 1

	var queue []*DisplayNode
	level := 0
	queue = append(queue, &DisplayNode{Thing: t.Root, Level: level})

	for len(queue) != 0 {
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
		}

		if node.Right != nil {
			newNode := &DisplayNode{
				Thing: node.Right,
				Level: level + 1,
			}
			queue = append(queue, newNode)
		}

		node.display(level, maxLevel)
	}

	//   6
	// /   \
	//3     8

	//     6
	//   /   \
	//  3     8
	// / \   / \
	//1   5 7   9

	//          6
	//        /   \
	//     3          8
	//   /  \        /  \
	//  1     5     7     9
	// / \   / \   / \   / \
	////x   x x   x x   x x   10

	// whiteRed.Println("HELLO WORLD\n\n\n")
	// whiteBlack.Println("HELLO WORLD\n\n\n")

	// //BFS and get color from struct
	// whiteBlack.Printf("%d\n", t.Root.Value)
	// whiteRed.Printf("%d\n", t.Root.Left.Value)
	// whiteRed.Printf("%d\n", t.Root.Right.Value)
}

// Handle nulls better
// What if the numbers are multiple digits...
func (n *Node) display(level, maxLevel int) {
	red := color.New(color.FgRed)
	black := color.New(color.FgBlack)
	whiteRed := red.Add(color.BgWhite)
	whiteBlack := black.Add(color.BgWhite)

	msg := ""
	numSpaces := int(math.Pow(2.0, float64(maxLevel)) - math.Pow(2.0, float64(level)))
	for i := 0; i < numSpaces; i++ {
		msg += " "
	}
	fmt.Print(msg)
	if n.Red {
		whiteRed.Printf("%d\n", n.Value)
	} else {
		whiteBlack.Printf("%d\n", n.Value)
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
