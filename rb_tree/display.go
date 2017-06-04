package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	t.formatForVisJS()

	bytes, err := json.Marshal(jsonNodes)
	writeJSONFile("nodes.json", bytes, err)

	bytes, err = json.Marshal(jsonEdges)
	writeJSONFile("edges.json", bytes, err)
}

func writeJSONFile(name string, bytes []byte, err error) {
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(name, bytes, 0644)
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
