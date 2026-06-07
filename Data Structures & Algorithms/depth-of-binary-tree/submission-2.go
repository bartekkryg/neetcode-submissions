/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type queue struct {
	values []any
}

func newQ() *queue {
	return &queue{
		values: make([]any, 0),
	}
}

func (q *queue) enqueue(val any) {
	q.values = append(q.values, val)
}

func (q *queue) dequeue() any {
	var val any
	if len(q.values) >= 2 {
		val = q.values[0]
		q.values = q.values[1:]
	} else if len(q.values) == 1  {
		 val = q.values[0]
		 q.values = make([]any, 0)
	} else if len(q.values) == 0 {
		val = nil
	}
	return val
}

func (q *queue) isEmpty() bool {
	return len(q.values) == 0 
}

func (q *queue) Len() int {
	return len(q.values)
}

func maxDepth(root *TreeNode) int {
    if root == nil {
		return 0
	}
	var level int
	q := newQ()
	q.enqueue(root)
	for !q.isEmpty() {
		for range q.Len(){
			node := (q.dequeue()).(*TreeNode)
			if node.Left != nil {
				q.enqueue(node.Left)
			}
			if node.Right != nil {
				q.enqueue(node.Right)
			}
		}
		level++
	}
	return level
}
