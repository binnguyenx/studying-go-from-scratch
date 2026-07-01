// Effective Go — Semicolons
// https://go.dev/doc/effective_go#semicolons
//
// Go inserts semicolons; braces required. Line breaks do not end statements like Python.

package gobasics

// NeedsBraces demonstrates mandatory braces (semicolons inserted by compiler).
func NeedsBraces(x int) int {
	sum := 0
	for i := 0; i < x; i++ {
		sum += i
	}
	return sum
}
