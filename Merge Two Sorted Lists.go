/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package leetcodetest

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	node := head
	if l1 == nil && l2 == nil {
		return nil
	}
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				node.Val = l1.Val
				l1 = l1.Next
			} else if l1.Val > l2.Val {
				node.Val = l2.Val
				l2 = l2.Next
			} else {
				node.Val = l1.Val
				node.Next = &ListNode{l2.Val, nil}
				node = node.Next
				l1 = l1.Next
				l2 = l2.Next
			}
		} else if l1 == nil || l2 == nil {
			if l1 == nil {
				node.Val = l2.Val
				l2 = l2.Next
			} else {
				node.Val = l1.Val
				l1 = l1.Next
			}
		}
		if l1 != nil || l2 != nil {
			node.Next = new(ListNode)
			node = node.Next
		}
	}
	return head
}
