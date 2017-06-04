package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Font struct {
	Color string `json:"color"`
}

type jsonNode struct {
	ID    int    `json:"id"`
	Color string `json:"color"`
	Label string `json:"label"`
	Font  *Font  `json:"font"`
	node  *Node
}

type jsonEdge struct {
	From int `json:"from"`
	To   int `json:"to"`
}

var jsonNodes []*jsonNode
var jsonEdges []*jsonEdge

func colorToString(n *Node) string {
	if n.Red {
		return "red"
	}
	return "black"
}

func (t *Tree) Display() {
	f, err := os.OpenFile("nodes.json", os.O_CREATE|os.O_WRONLY, 777)
	if err != nil {
		panic("Can't write to nodes.json")
	}

	t.formatForVisJS()
	bytes, err := json.Marshal(jsonNodes)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(f, string(bytes))

	f, err = os.OpenFile("edges.json", os.O_CREATE|os.O_WRONLY, 777)
	if err != nil {
		panic("Can't write to edges.json")
	}

	bytes, err = json.Marshal(jsonEdges)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(f, string(bytes))
}

func (t *Tree) formatForVisJS() {
	newNode := &jsonNode{
		ID:    t.Root.Value,
		Color: colorToString(t.Root),
		Label: fmt.Sprintf("Node: %d", t.Root.Value),
		Font:  &Font{Color: "white"},
		node:  t.Root,
	}
	queue := []*jsonNode{newNode}
	jsonNodes = append(jsonNodes, newNode)

	for len(queue) != 0 {
		element := queue[0]
		queue = queue[1:]

		node := element.node

		if node.Left != nil {
			newNode := &jsonNode{
				ID:    node.Left.Value,
				Color: colorToString(node.Left),
				Label: fmt.Sprintf("Node: %d", node.Left.Value),
				Font:  &Font{Color: "white"},
				node:  node.Left,
			}
			newEdge := &jsonEdge{
				From: node.Value,
				To:   node.Left.Value,
			}
			jsonEdges = append(jsonEdges, newEdge)
			queue = append(queue, newNode)
			jsonNodes = append(jsonNodes, newNode)
		}
		if node.Right != nil {
			newNode := &jsonNode{
				ID:    node.Right.Value,
				Color: colorToString(node.Right),
				Label: fmt.Sprintf("Node: %d", node.Right.Value),
				Font:  &Font{Color: "white"},
				node:  node.Right,
			}
			newEdge := &jsonEdge{
				From: node.Value,
				To:   node.Right.Value,
			}
			jsonEdges = append(jsonEdges, newEdge)
			queue = append(queue, newNode)
			jsonNodes = append(jsonNodes, newNode)
		}
	}
}
