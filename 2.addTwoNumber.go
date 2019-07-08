package leetcode_go

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	 Val int
	 Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sum, carry int
	var head = &ListNode{0, nil }
	current := head
	carry = 0
	for l1 != nil && l2 != nil  {
		sum = l1.Val + l2.Val
		current.Next = &ListNode{
			Val: sum%10 + carry,
		}
		carry = sum/10
		current = current.Next
		l1, l2 = l1.Next, l2.Next
	}
	if l1 != nil {
		current.Next = l1
	}
	if l2 != nil {
		current.Next = l2
	}
	for carry == 1 {
		if current.Next == nil {
			current.Next = &ListNode{carry, nil }
			carry = 0
			return head.Next
		}
		val := current.Val + carry
		current.Val = val%10;
		carry = val/10
		current = current.Next;
	}
	return head.Next
}