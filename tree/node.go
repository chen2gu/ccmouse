package main

import (
	"fmt"
)

type treeNode struct {
	value       int
	left, right *treeNode
}

func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func (node treeNode) print() {
	fmt.Print(node.value, "  ")
}
func (node *treeNode) setValue(value int) {
	node.value = value
}
func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)

	root.print()
	root.right.left.setValue(4)
	fmt.Println()
	root.right.left.print()
	fmt.Println()

	root.print()
	root.setValue(100)

	fmt.Print(root)
	fmt.Println()

	pRoot := &root
	pRoot.print()
	pRoot.setValue(200)
	pRoot.print()

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)
}
