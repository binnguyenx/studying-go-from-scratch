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
// Khác khi so sánh == nil, JSON (null vs []), hoặc một số API phân biện “chưa có” vs “rỗng”.
//
// # Loop slice
//
//	for i := 0; i < len(numbers); i++ { fmt.Println(numbers[i]) }
//	for index, value := range numbers { fmt.Println(index, value) }
//
// Lưu ý range: value là BẢN COPY phần tử (với int/string OK; struct lớn tốn copy).
//
//	for _, value := range numbers { value = 100 }  // KHÔNG đổi slice
//	for i := range numbers { numbers[i] = 100 }   // đúng — sửa qua index
//
// Với struct trong slice, muốn sửa field: numbers[i].Name = "x" hoặc dùng []*T.
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

// SumSliceIndex — loop kiểu C: for i := 0; i < len(s); i++.
func SumSliceIndex(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

// SumSliceRange — for index, value := range nums.
func SumSliceRange(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

// FillSliceWrong — gán value trong range không ảnh hưởng slice gốc.
func FillSliceWrong(nums []int, v int) {
	for _, n := range nums {
		n = v // n là copy; slice không đổi
		_ = n
	}
}

// FillSliceCorrect — sửa qua index (hoặc for i := 0; i < len; i++).
func FillSliceCorrect(nums []int, v int) {
	for i := range nums {
		nums[i] = v
	}
}
