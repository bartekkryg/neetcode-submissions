/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 * 
 */

func invertTree(root *TreeNode) *TreeNode {
	DFS(root)
	return root
}

func swap(node *TreeNode){
	left := node.Left
	node.Left = node.Right
	node.Right = left
}

func DFS(node *TreeNode) {
	if node == nil {
		return 
	}
	swap(node)
	DFS(node.Left)
	DFS(node.Right)
}