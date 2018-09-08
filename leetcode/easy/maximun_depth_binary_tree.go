package easy

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left:= maxDepth(root.Left)
	right := maxDepth(root.Right)

	return int(math.Max(float64(left), float64(right)))+1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}






