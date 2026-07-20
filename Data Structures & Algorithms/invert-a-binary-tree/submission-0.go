/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
		return nil
	}

	// swap left and right children
	root.Left, root.Right = root.Right, root.Left

	// recursively invert left and right subtrees
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}
