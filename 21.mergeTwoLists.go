package leetcode_go

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for l1 != nil || l2 != nil  {
		if l1 == nil {
			current.Next = l2
			l2 = nil
		} else if l2 == nil {
			current.Next = l1
			l1 = nil
		} else {
			l1Val, l2Val, Val := l1.Val, l2.Val, 0;
			if l1Val < l2Val {
				Val = l1Val
				l1 = l1.Next
			} else {
				Val = l2Val
				l2 = l2.Next
			}
			current.Next = &ListNode{ Val:Val }
			current = current.Next
		}
	}
	return dummy.Next
}