/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Next: head}

	slow := dummy
	fast := dummy

	// move fast n+1 steps ahead
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	// move both the pointers together
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}
