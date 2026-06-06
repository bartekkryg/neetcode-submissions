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
	// var lD, rD int
	// lD += DFS(root.Left)
	// rD += DFS(root.Right)
    return DFS(root)
}

func DFS(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + max(DFS(node.Left), DFS(node.Right))
}

func max(a, b int)int {
	if a >b {
		return a
	}
	return b
}