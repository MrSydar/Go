package shamir

import (
	"reflect"
	"testing"
)

func TestGetParts(t *testing.T) {
	pN := 4
	pA := []int{954, 352, 62}
	pP := 1523

	actual, err := GetParts(pA, pN, pP)

	if err != nil {
		t.Fatalf("error is not expected: %v", err)
	}

	if len(actual) != pN {
		t.Fatalf("expected parts length to be %v, but got %v", pN, len(actual))
	}

	expected := []Part{
		{1368, 1},
		{383, 2},
		{1045, 3},
		{308, 4},
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected parts to be %v, but got %v", expected, actual)
	}
}

func TestGetSecret(t *testing.T) {
	selectedParts := []Part{
		{1368, 1},
		{1045, 3},
		{308, 4},
	}

	expected := 954
	actual, err := GetSecret(selectedParts, 3, 1523)

	if err != nil {
		t.Fatal("unexpected error")
	}

	if actual != expected {
		t.Fatalf("expected secret to be %v, but got %v", expected, actual)
	}
}
