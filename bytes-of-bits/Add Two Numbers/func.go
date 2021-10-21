package addtwonumbers

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func doOneDigitSum(a int, b int) (sum int, carry bool) {
	sum = a + b

	if sum >= 10 {
		return sum - 10, true
	} else {
		return sum, false
	}
}

func carryAdd(l *ListNode, prevCarry bool) (solution *ListNode) {
	if prevCarry {
		return zipAdd(l, &ListNode{1, nil}, false)
	} else {
		return l
	}
}

func zipAdd(l1 *ListNode, l2 *ListNode, prevCarry bool) (solution *ListNode) {
	if l1 == nil {
		return carryAdd(l2, prevCarry)
	} else if l2 == nil {
		return carryAdd(l1, prevCarry)
	} else {
		var sum int
		var carry bool

		if prevCarry {
			sum, carry = doOneDigitSum(l1.Val, l2.Val+1)
		} else {
			sum, carry = doOneDigitSum(l1.Val, l2.Val)
		}

		return &ListNode{sum, zipAdd(l1.Next, l2.Next, carry)}
	}
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return zipAdd(l1, l2, false)
}

func ArrayToListNode(arr []int) *ListNode {
	if len(arr) == 1 {
		return &ListNode{arr[0], nil}
	} else {
		return &ListNode{arr[0], ArrayToListNode(arr[1:])}
	}
}

func (l1 *ListNode) Equals(l2 *ListNode) bool {
	if l1 == nil && l2 != nil || l1 != nil && l2 == nil || l1 != nil && l2 != nil && l1.Val != l2.Val {
		return false
	} else if l1 == nil && l2 == nil {
		return true
	} else {
		return l1.Next.Equals(l2.Next)
	}
}

func (l *ListNode) ToString() string {
	if l.Next != nil {
		return fmt.Sprintf("%d, %s", l.Val, l.Next.ToString())
	} else {
		return fmt.Sprint(l.Val)
	}
}
