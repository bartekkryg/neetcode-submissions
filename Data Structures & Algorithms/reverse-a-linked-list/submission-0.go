/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	node := head // 0
	var prev *ListNode
	for node != nil  {
		tmp := node.Next // 1, 1.next = 2
		node.Next = prev // 0.next = prev
		prev = node // prev = 0
		node = tmp // node = 1
	}
	return prev
}
