package main

import (
	"fmt"

	"github.com/fatih/color"
)

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
