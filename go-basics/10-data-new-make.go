// Effective Go — new & make
// https://go.dev/doc/effective_go#allocation_new
// https://go.dev/doc/effective_go#allocation_make
//
// new(T) → *T, zeroed. make(slice|map|chan, ...) → initialized value (not pointer).

package gobasics

// NewIntPointer — new(T) returns *T zeroed.
func NewIntPointer() *int {
	return new(int)
}

// MakeIntSlice — make([]T, len, cap); initialized, ready to use.
func MakeIntSlice() []int {
	return make([]int, 10, 100)
}

// NewVsMakeNilSlice — new([]int) returns *slice header pointing at nil slice (rare).
func NewVsMakeNilSlice() (*[]int, []int) {
	p := new([]int)
	v := make([]int, 10)
	return p, v
}
