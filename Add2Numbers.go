package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	l2 := &ListNode{2, &ListNode{4, &ListNode{6, nil}}}
	l3 := addTwoNumbers(l1, l2)
	for l3 != nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	T3 := new(ListNode)
	node := T3
	sum := 0
	for l1 != nil || l2 != nil || sum != 0 {
		num := 0
		if l1 != nil {
			num += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			num += l2.Val
			l2 = l2.Next
		}
		sum += num
		node.Next = &ListNode{sum % 10, nil}
		sum = sum / 10
		node = node.Next
	}
	return T3.Next
}
