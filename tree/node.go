package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func (node treeNode) print() {
	fmt.Println(node.value)
}
func (node *treeNode) setValue(value int) {
	node.value = value
}

func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.print()
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = createNode(3)
	root.print()
	root.setValue(5)
	root.print()

}

func createNode(value int) *treeNode {
	return &treeNode{value: value}
}
