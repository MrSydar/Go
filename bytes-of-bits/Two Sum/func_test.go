package twosum

import (
	"reflect"
	"testing"
)

func TestTwoFunc(t *testing.T) {
	arr := make([]int, 5)
	arr[0], arr[1], arr[2], arr[3], arr[4] = 3, 1, 4, 6, 7

	expected := make([]int, 2)
	expected[0], expected[1] = 0, 4

	actual := TwoSum(arr, 10)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %v but got %v", expected, actual)
	}
}
