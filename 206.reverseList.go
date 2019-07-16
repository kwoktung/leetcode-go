package leetcode_go


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		node := head
		head = head.Next
		node.Next = prev
		prev = node
	}
	return prev
}