// Effective Go — If & redeclaration
// https://go.dev/doc/effective_go#if
// https://go.dev/doc/effective_go#redeclaration
//
// No parens around condition; braces mandatory.
// if err := f(); err != nil { return err }
// := can reassign err if at least one new name in same scope.
package gobasics

import "errors"

var errRedecl = errors.New("step")

func Sign(n int) int {
	if n < 0 {
		return -1
	}
	if n == 0 {
		return 0
	}
	return 1
}

func ErrRedeclarationDemo() (int, error) {
	first, err := step(1)
	if err != nil {
		return 0, err
	}
	second, err := step(2)
	if err != nil {
		return first, err
	}
	return first + second, nil
}

func step(v int) (int, error) {
	if v < 0 {
		return 0, errRedecl
	}
	return v, nil
}
