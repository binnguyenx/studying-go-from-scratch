// Effective Go — pointers.go
// https://go.dev/doc/effective_go
//
package gobasics

func Swap(a, b *int) {
	if a == nil || b == nil {
		return
	}
	*a, *b = *b, *a
}

func PtrInt(v int) *int {
	return &v
}

func Deref(p *int, fallback int) int {
	if p == nil {
		return fallback
	}
	return *p
}
