package main

import (
	"fmt"
	"github.com/chen2guo/ccmouse/tree"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()

	right.postOrder()
	myNode.node.Print()

}

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CeateNode(2)
	root.Right.Left.SetValue(4)

	root.Traverse()

	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node Count: ", nodeCount)

	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()

	fmt.Println("------")
	root.Print()
	root.SetValue(100)

	fmt.Print(root)
	fmt.Println()

	pRoot := &root
	pRoot.Print()
	pRoot.SetValue(200)
	pRoot.Print()

	var pROOT *tree.Node
	pROOT.SetValue(200)
	pROOT = &root
	pROOT.SetValue(300)
	pROOT.Print()

	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

}
