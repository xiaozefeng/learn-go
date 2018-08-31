package main

import "fmt"

func main() {
	// init data
	root := Tree{Value: 3}
	root.Left = &Tree{Value: 1}
	root.Left.Left = &Tree{Value: 1}
	root.Left.Right = &Tree{Value: 2}

	root.Right = &Tree{Value: 8}
	root.Right.Left = &Tree{Value: 5}
	root.Right.Right = &Tree{Value: 13}

	//ch := root.traverseWithChannel()
	//for v := range ch {
	//	fmt.Printf("%d ", v.Value)
	//}

	ch := make(chan int)
	walk(&root, ch)
	for v := range ch{
		fmt.Printf("%d ", v)
	}

}

type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

func (t *Tree) traverse() {
	t.traverseFunc(func(tree *Tree) {
		fmt.Printf("%d ", tree.Value)
	})
}

func (t *Tree) traverseFunc(f func(tree *Tree)) {
	if t == nil {
		return
	}
	t.Left.traverseFunc(f)
	f(t)
	t.Right.traverseFunc(f)
}

func (t *Tree) traverseWithChannel() chan *Tree {
	ch := make(chan *Tree)
	go func() {
		t.traverseFunc(func(tree *Tree) {
			ch <- tree
		})
		close(ch)
	}()
	return ch
}

func walk(root *Tree, ch chan int) {
	go func() {
		root.traverseFunc(func(tree *Tree) {
			ch <- tree.Value
		})
		close(ch)
	}()
}

func same(t1, t2 *Tree) bool {
	return t1.Value == t2.Value
}
