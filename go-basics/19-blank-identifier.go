// Effective Go — The blank identifier
// https://go.dev/doc/effective_go#blank
//
// _ discards a value: for _, v := range xs { }, import _ "pkg" for side effects.

package gobasics

import "errors"

// SumValues uses _ to ignore index in range.
func SumValues(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

var errBlankDemo = errors.New("demo")

// MultiAssignBlank — second := reassigns err (Effective Go redeclaration).
func MultiAssignBlank() error {
	_, err := fail(true)
	if err != nil {
		return err
	}
	_, err = fail(false)
	return err
}

func fail(f bool) (int, error) {
	if f {
		return 0, errBlankDemo
	}
	return 1, nil
}
