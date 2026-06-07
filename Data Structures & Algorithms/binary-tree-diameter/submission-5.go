/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func diameterOfBinaryTree(root *TreeNode) int {
	d := 0
	if root == nil {
		return 0
	}
	var DFS func(*TreeNode) int
	DFS = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := DFS(node.Left)
		r := DFS(node.Right)
		d = max(d, l+r)
		return 1 + max(l, r)
	}
	DFS(root)
    return d
}
