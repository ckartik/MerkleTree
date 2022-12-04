package main

import (
	"fmt"
)

type MerkleNode struct {
	isLeaf     bool
	leftChild  *MerkleNode
	rightChild *MerkleNode
	payload    [64]byte
}

func main() {
	fmt.Println("hello")
}
