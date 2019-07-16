package leetcode_go

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next:head}
	prev := dummy
	for prev.Next != nil && prev.Next.Next != nil  {
		a, b := prev.Next, prev.Next.Next
		prev.Next = b
		a.Next = b.Next
		b.Next = a
		prev = a
	}
	return dummy.Next
}
