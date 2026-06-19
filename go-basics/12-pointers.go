// Pointers — & address, * dereference, nil
//
// p := &x
// *p = 10
// func swap(a, b *int)
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
