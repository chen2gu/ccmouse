package tree

import (
	"fmt"
)

type Node struct {
	Value       int
	Left, Right *Node
}

func CeateNode(value int) *Node {
	return &Node{Value: value}
}

func (node Node) Print() {
	fmt.Print(node.Value, "  ")
}
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored. ")
		return
	}
	node.Value = value
}
