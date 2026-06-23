// Slices — dynamic view []T, len, cap, append, copy, shared backing array
//
// # Tạo slice
//   s := []int{1,2,3}
//   s := make([]int, len, cap)
//   append(s, 4)
//   copy(dst, src)
//
// # nil slice vs empty slice (len đều = 0, nhưng khác nhau)
//
//   var nilSlice []int     // nil — chưa trỏ backing array; nilSlice == nil → true
//   empty := []int{}       // empty — đã allocate header; empty == nil → false
//   empty2 := make([]int, 0) // empty tương tự literal {}
//
// Cả hai đều len=0, cap=0 (thường), và append đều OK.
// Khác khi so sánh == nil, JSON (null vs []), hoặc một số API phân biệt “chưa có” vs “rỗng”.
package gobasics

func SliceAppend(nums []int, more ...int) []int {
	return append(nums, more...)
}

func ReverseInPlace(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func SubSlice(s []int, low, high int) []int {
	return s[low:high]
}

// NilSlice — var s []int; zero value của slice.
func NilSlice() []int {
	var s []int
	return s
}

// EmptySliceLiteral — []T{}; len 0 nhưng không nil.
func EmptySliceLiteral() []int {
	return []int{}
}

// EmptySliceMake — make([]T, 0); empty, có thể truyền cap thứ 3: make([]int, 0, 8).
func EmptySliceMake() []int {
	return make([]int, 0)
}

// EmptySliceMakeCap pre-allocates backing capacity while len stays 0.
func EmptySliceMakeCap(capacity int) []int {
	return make([]int, 0, capacity)
}

// SliceIsNil reports whether s is a nil slice (not the same as len(s)==0).
func SliceIsNil(s []int) bool {
	return s == nil
}

// SliceLenCap returns len and cap (both 0 for nil/empty unless cap pre-allocated).
func SliceLenCap(s []int) (length, capacity int) {
	return len(s), cap(s)
}

// AppendToEither works on nil or empty slice — append allocates on first use.
func AppendToEither(s []int, v int) []int {
	return append(s, v)
}
