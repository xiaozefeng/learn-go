package easy

import (
	"testing"
	"fmt"
)

func TestMergeTwoLists(t *testing.T) {
	//tests := []struct {
	//	l1 *ListNode
	//	l2 *ListNode
	//	l3 *ListNode
	//}{
	//	{
	//		&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}},
	//		&ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
	//		&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}},
	//	},
	//}
	//
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	l1.traver()
	fmt.Println()
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	l2.traver()
	fmt.Println()
	listNode := MergeTwoLists(l1, l2)
	listNode.traver()
	fmt.Println()
	
}
