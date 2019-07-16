package leetcode_go

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next:head}
	start, end := head, dummy
	for n > 0 && start != nil {
		start = start.Next
		n--
	}
	if start == nil {
		return head.Next
	}
	for start != nil {
		start = start.Next;
		end = end.Next
	}
	end.Next = end.Next.Next
	return head
}


func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	i, m, current := 0, map[int]*ListNode{}, head
	for current != nil  {
		m[i] = current
		current = current.Next
		i++
	}

	l := len(m)
	target := l - n
	if node, ok := m[target - 1]; ok {
		node.Next = m[target].Next
		m[target].Next = nil
	} else {
		head = m[target].Next
	}
	return head
}