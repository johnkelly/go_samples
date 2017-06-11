package main

import (
	"encoding/json"
	"io/ioutil"
)

type Font struct {
	Color string `json:"color"`
}

type jsonNode struct {
	ID       string `json:"id"`
	Color    string `json:"color"`
	Label    string `json:"label"`
	Font     *Font  `json:"font"`
	node     *Node
	parentID string
}

type jsonEdge struct {
	From string `json:"from"`
	To   string `json:"to"`
}

var jsonNodes []*jsonNode
var jsonEdges []*jsonEdge

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
		ID:    t.root.hash,
		Color: "green",
		Label: t.root.hash[:6],
		Font:  &Font{Color: "white"},
		node:  t.root,
	}
	queue := []*jsonNode{newNode}
	jsonNodes = append(jsonNodes, newNode)

	for len(queue) != 0 {
		element := queue[0]
		queue = queue[1:]

		node := element.node

		if element.parentID != "" {
			newEdge := &jsonEdge{
				From: element.parentID,
				To:   element.ID,
			}
			jsonEdges = append(jsonEdges, newEdge)
		}

		if node.left != nil {
			id := node.left.hash + "left"

			newNode := &jsonNode{
				ID:       id,
				Color:    "green",
				Label:    node.left.hash[:6],
				Font:     &Font{Color: "white"},
				node:     node.left,
				parentID: element.ID,
			}
			queue = append(queue, newNode)
			jsonNodes = append(jsonNodes, newNode)
		}
		if node.right != nil {
			id := node.right.hash + "right"

			newNode := &jsonNode{
				ID:       id,
				Color:    "green",
				Label:    node.right.hash[:6],
				Font:     &Font{Color: "white"},
				node:     node.right,
				parentID: element.ID,
			}
			queue = append(queue, newNode)
			jsonNodes = append(jsonNodes, newNode)
		}
	}
}
