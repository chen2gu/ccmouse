package main

import (
	"fmt"
	"github.com/chen2gu/ccmouse/tree"
)

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CeateNode(2)

	root.Print()
	root.Right.Left.SetValue(4)
	fmt.Println()
	root.Right.Left.Print()
	fmt.Println()

	root.Traverse()

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
