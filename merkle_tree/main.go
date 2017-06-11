package main

import (
	"crypto/sha256"
	"fmt"
)

type Node struct {
	left  *Node
	right *Node
	hash  string
}

type Tree struct {
	root *Node
}

func main() {
	values := []string{"apple", "banana", "carrot", "date", "eggplant", "fennel", "goat", "habanero", "ice"}
	hashes := generateHashes(values)
	nodes := generateNodes(hashes)

	tree := &Tree{root: buildMerkleTree(nodes)}

	fmt.Printf("Merkle Root:\t%v\n", tree.root.hash)
	tree.Display()
}

func buildMerkleTree(nodes []*Node) *Node {
	for len(nodes) > 1 {
		newNodes := []*Node{}

		if len(nodes)%2 != 0 {
			lastNode := nodes[len(nodes)-1]
			node := &Node{hash: lastNode.hash}
			nodes = append(nodes, node)
		}
		for i := 0; i < len(nodes); i = i + 2 {
			newNode := &Node{
				left:  nodes[i],
				right: nodes[i+1],
				hash:  combinedHash(nodes[i].hash, nodes[i+1].hash),
			}

			newNodes = append(newNodes, newNode)
		}
		nodes = newNodes
	}
	return nodes[0]
}

func combinedHash(hashA, hashB string) string {
	combined := hashA + hashB
	hash := sha256.Sum256([]byte(combined))
	return fmt.Sprintf("%X", hash)
}

func generateNodes(hashes []string) (nodes []*Node) {
	for _, hash := range hashes {
		nodes = append(nodes, &Node{hash: hash})
	}
	return nodes
}

func generateHashes(values []string) (hashes []string) {
	for _, value := range values {
		hash := sha256.Sum256([]byte(value))
		hexDigest := fmt.Sprintf("%X", hash)
		hashes = append(hashes, hexDigest)
	}
	return hashes
}
