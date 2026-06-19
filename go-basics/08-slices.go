// Slices — dynamic view []T, len, cap, append, copy, shared backing array
//
// s := []int{1,2,3}
// s := make([]int, len, cap)
// append(s, 4)
// copy(dst, src)
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
