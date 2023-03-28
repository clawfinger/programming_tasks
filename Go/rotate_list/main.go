package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func length(head *ListNode) int {
	current := head
	res := 0
	for current != nil {
		res++
		current = current.Next
	}
	return res
}

func kthElement(head *ListNode, k int) (*ListNode, *ListNode) {
	current := head
	next := head
	for i := 0; i < k; i++ {
		current = next
		next = current.Next
	}
	return current, next
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	len := length(head)
	if len == 1 {
		return head
	}
	trueK := k % len
	if trueK == 0 {
		return head
	}
	firstEnd, secondStart := kthElement(head, len-trueK)
	firstEnd.Next = nil
	secondEnd, _ := kthElement(secondStart, trueK)
	secondEnd.Next = head
	return secondStart
}
