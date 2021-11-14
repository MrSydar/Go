package trivial

import (
	"reflect"
	"testing"
)

func TestGenerateTrivialPartsForKEq1000AndNEq3AndSRandom(t *testing.T) {
	n := 3
	k := 1000

	_, err := GetPartsRandom(n, k)
	if err != nil {
		t.Fatalf("error was not expected")
	}
}

func TestGenerateTrivialPartsForKEq1000AndNEq3AndSEq456AndS1Eq856AndS2Eq231(t *testing.T) {
	arr := []int{456, 856, 231}
	k := 1000

	actual, _ := GetParts(arr, k)
	if expected := []int{856, 231, 369}; !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected parts %v, but got %v", expected, actual)
	}
}

func TestGenerateTrivialPartsOneElementEqualsK(t *testing.T) {
	arr := []int{456, 1000, 231}
	k := 1000

	actual, err := GetParts(arr, k)

	if actual != nil {
		t.Fatalf("expected parts to be nil, but got %v", actual)
	}

	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestGenerateTrivialPartsOneElementGreaterThanK(t *testing.T) {
	arr := []int{456, 1005, 231}
	k := 1000

	actual, err := GetParts(arr, k)

	if actual != nil {
		t.Fatalf("expected parts to be nil, but got %v", actual)
	}

	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestGenerateTrivialPartsKEq0(t *testing.T) {
	arr := []int{456, 856, 231}
	k := 0

	actual, err := GetParts(arr, k)

	if actual != nil {
		t.Fatalf("expected parts to be nil, but got %v", actual)
	}

	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestGenerateTrivialSecretForKEq1000AndNEq3AndS1Eq856AndS2Eq231AndS3Eq369(t *testing.T) {
	parts := []int{856, 231, 369}
	k := 1000

	actual, _ := GetSecret(parts, k)
	if expected := 456; actual != expected {
		t.Fatalf("expected secret %d, but got %d", expected, actual)
	}
}

func TestGenerateTrivialSecretOneElementEqualsK(t *testing.T) {
	parts := []int{856, 1000, 269}
	k := 1000

	actual, err := GetSecret(parts, k)

	if actual != -1 {
		t.Fatalf("expected secret to be -1, but got %v", actual)
	}

	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestGenerateTrivialSecretOneElementGreaterThanK(t *testing.T) {
	parts := []int{856, 1005, 269}
	k := 1000

	actual, err := GetSecret(parts, k)

	if actual != -1 {
		t.Fatalf("expected secret to be -1, but got %v", actual)
	}

	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestGenerateTrivialSecretKEq0(t *testing.T) {
	parts := []int{856, 231, 269}
	k := 0

	actual, err := GetSecret(parts, k)

	if actual != -1 {
		t.Fatalf("expected secret to be -1, but got %v", actual)
	}

	if err == nil {
		t.Fatalf("expected error")
	}
}
