package main

import (
	"fmt"
	"imooc.com/learn-go/tree"
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
	root.Print()
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = tree.CreateNode(3)
	root.Print()
	root.SetValue(4)
	root.Print()

	fmt.Println("遍历tree node")
	root.Traverse()

	fmt.Println("------------")
	n := &myTreeNode{&root}
	n.postOrder()

	fmt.Println("--- 统计有多少个元素")
	nodeCount := 0
	root.TraverseFunc(func(n *tree.Node) {
		nodeCount ++
	})
	fmt.Println(nodeCount)

	fmt.Println("使用channel遍历节点")
	c := root.TraverseWithChannel()
	maxNode :=0
	for node:= range c{
		if node.Value > maxNode{
			maxNode = node.Value
		}
	}
	fmt.Println("max node value: ", maxNode)
}
