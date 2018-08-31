package easy

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	return nil
}

func (n *ListNode) traver() {
	if n == nil {
		return
	}
	fmt.Printf("%d ", n.Val)
	n.Next.traver()
}
