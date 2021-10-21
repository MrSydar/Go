package addtwonumbers

import (
	"testing"
)

func TestEqualsSameMultipleElementsTillEndOfShorterNotSameLength(t *testing.T) {
	l1 := &ListNode{0, &ListNode{5, &ListNode{-1, nil}}}
	l2 := &ListNode{0, &ListNode{5, &ListNode{-1, &ListNode{-1, nil}}}}

	if l1.Equals(l2) || l2.Equals(l1) {
		t.Fatal("Lists are not equal")
	}
}

func TestEqualsNotSameMultipleElementsSameLength(t *testing.T) {
	l1 := &ListNode{0, &ListNode{1, &ListNode{-1, nil}}}
	l2 := &ListNode{0, &ListNode{5, &ListNode{-1, nil}}}

	if l1.Equals(l2) || l2.Equals(l1) {
		t.Fatal("Lists are not equal")
	}
}

func TestEqualsSameMultipleElementsSameLength(t *testing.T) {
	l1 := &ListNode{0, &ListNode{1, &ListNode{-4, nil}}}
	l2 := &ListNode{0, &ListNode{1, &ListNode{-4, nil}}}

	if !l1.Equals(l2) && !l2.Equals(l1) {
		t.Fatal("Lists are equal")
	}
}

func TestEqualsNotSameSingleElement(t *testing.T) {
	l1 := &ListNode{0, nil}
	l2 := &ListNode{1, nil}

	if l1.Equals(l2) || l2.Equals(l1) {
		t.Fatal("Lists are not equal")
	}
}

func TestEqualsSameSingleElement(t *testing.T) {
	l1 := &ListNode{0, nil}
	l2 := &ListNode{0, nil}

	if !l1.Equals(l2) && !l2.Equals(l1) {
		t.Fatal("Lists are equal")
	}
}

func TestArrayToListNode(t *testing.T) {
	l1 := &ListNode{0, &ListNode{1, &ListNode{2, nil}}}
	l2 := ArrayToListNode([]int{0, 1, 2})

	if !l1.Equals(l2) {
		t.Fatal("Bad list node generated")
	}
}

func TestToString(t *testing.T) {
	l := ArrayToListNode([]int{0, 1, 2})
	lStr := "0, 1, 2"

	if l.ToString() != lStr {
		t.Fatal("Bad string generated")
	}
}

func TestAddTwoSmallNumbers(t *testing.T) {
	l1 := ArrayToListNode([]int{1, 2, 3})
	l2 := ArrayToListNode([]int{2, 1, 3})
	expected := ArrayToListNode([]int{3, 3, 6})
	actual := AddTwoNumbers(l1, l2)

	if !actual.Equals(expected) {
		t.Fatalf("Bad sum. Expected %s, but got %s.", expected.ToString(), actual.ToString())
	}
}

func TestAddLargeWithCarryNumbers(t *testing.T) {
	l1 := ArrayToListNode([]int{9, 9, 9, 9})
	l2 := ArrayToListNode([]int{9, 9, 9, 9})
	expected := ArrayToListNode([]int{8, 9, 9, 9, 1})
	actual := AddTwoNumbers(l1, l2)

	if !actual.Equals(expected) {
		t.Fatalf("Bad sum. Expected %s, but got %s.", expected.ToString(), actual.ToString())
	}
}
