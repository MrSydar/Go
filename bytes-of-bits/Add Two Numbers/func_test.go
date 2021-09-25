package addtwonumbers

import (
	"testing"
)

func arrayToListNode(arr []int) *ListNode {
	if len(arr) == 1 {
		return &ListNode{arr[0], nil}
	} else {
		return &ListNode{arr[0], arrayToListNode(arr[1:])}
	}
}

func (l1 *ListNode) equals(l2 *ListNode) bool {
	if l1 == nil && l2 != nil || l1 != nil && l2 == nil || l1 != nil && l2 != nil && l1.Val != l2.Val {
		return false
	} else if l1 == nil && l2 == nil {
		return true
	} else {
		return l1.Next.equals(l2.Next)
	}
}

func TestAddTwoNumbers(t *testing.T) {
	l1, l2 := arrayToListNode([]int{2, 4, 3}), arrayToListNode([]int{5, 6, 4})
	expected1 := arrayToListNode([]int{7, 0, 8})
	if !expected1.equals(AddTwoNumbers(l1, l2)) {
		t.Fatalf("TC1")
	}

	l3, l4 := arrayToListNode([]int{0}), arrayToListNode([]int{0})
	expected2 := arrayToListNode([]int{0})
	if !expected2.equals(AddTwoNumbers(l3, l4)) {
		t.Fatalf("TC2")
	}

	l5, l6 := arrayToListNode([]int{9, 9, 9, 9, 9, 9, 9}), arrayToListNode([]int{9, 9, 9, 9})
	expected3 := arrayToListNode([]int{8, 9, 9, 9, 0, 0, 0, 1})
	if !expected3.equals(AddTwoNumbers(l5, l6)) {
		t.Fatalf("TC2")
	}
}
