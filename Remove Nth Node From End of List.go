package leetcodetest

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	s := size(head)
	n = s - n
	if n == 0 {
		return head.Next
	}
	h := head
	for i := 0; i < n-1; i++ {
		h = h.Next
	}
	h.Next = h.Next.Next
	return head
}

func size(head *ListNode) int {
	size := 0
	for h := head; h != nil; h = h.Next {
		size++
	}
	return size
}
