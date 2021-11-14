package trivial

import (
	"errors"
	"fmt"
)

func checkInput(s []int, k int) error {
	if k == 0 {
		return errors.New("the k paramether can't be zero")
	} else {
		for i := 0; i < len(s); i++ {
			if s[i] >= k {
				return fmt.Errorf("element on index %v can't be greater or equal the k value", i)
			}
		}
	}
	return nil
}

func GetPartsRandom(n int, k int) (parts []int, err error) {
	return nil, nil
}

func GetParts(s []int, k int) (parts []int, err error) {
	if iErr := checkInput(s, k); iErr != nil {
		return nil, iErr
	}

	diff := s[0]
	for _, val := range s[1:] {
		diff -= val
	}

	lastPart := ((diff % k) + k) % k

	return append(s[1:], lastPart), nil
}

func GetSecret(s []int, k int) (secret int, err error) {
	if iErr := checkInput(s, k); iErr != nil {
		return -1, iErr
	}

	sum := s[0]
	for _, val := range s[1:] {
		sum += val
	}

	return ((sum % k) + k) % k, nil
}
