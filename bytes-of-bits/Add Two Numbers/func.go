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

func sumFromEnd(l *ListNode, thisCh chan int, otherCh chan int, digitNum int, isPrimary bool) (solution *ListNode, carry bool) {
	digitNum += 1

	if l.Next == nil {
		thisCh <- digitNum
		otherDigitNum := <-otherCh

		if digitNum > otherDigitNum || digitNum == otherDigitNum && isPrimary {
			currDigitSum, currHasCarry := doOneDigitSum(l.Val, <-otherCh)

			return &ListNode{currDigitSum, nil}, currHasCarry
		} else {
			thisCh <- l.Val
			return
		}
	} else {
		lSum, prevHasCarry := sumFromEnd(l.Next, thisCh, otherCh, digitNum, isPrimary)

		if lSum != nil {
			var currDigitSum int
			var currHasCarry bool

			otherDigit := <-otherCh

			if prevHasCarry {
				currDigitSum, currHasCarry = doOneDigitSum(l.Val+1, otherDigit)
			} else {
				currDigitSum, currHasCarry = doOneDigitSum(l.Val, otherDigit)
			}

			return &ListNode{currDigitSum, lSum}, currHasCarry
		} else {
			thisCh <- l.Val
			return
		}
	}
}

func startSumFromEnd(l *ListNode, thisCh chan int, otherCh chan int, isPrimary bool, solution chan *ListNode) {
	lSum, prevHasCarry := sumFromEnd(l, thisCh, otherCh, 0, isPrimary)

	if lSum != nil {
		if prevHasCarry {
			lSum = &ListNode{1, lSum}
		}
		solution <- lSum
	} else {
		close(thisCh)
	}
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Ch, l2Ch := make(chan int, 1), make(chan int, 1)
	solution := make(chan *ListNode)

	go startSumFromEnd(l1, l1Ch, l2Ch, true, solution)
	go startSumFromEnd(l2, l2Ch, l1Ch, false, solution)

	return <-solution
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
