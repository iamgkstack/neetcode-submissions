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

	// find Left & right depths
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	// return the larger depth + current node
	if left > right {
		return left+1
	}

	return right+1
}
