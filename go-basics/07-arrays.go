// Arrays — fixed length [N]T, value type (copied on assign)
//
// # Khai tạo array
//
//	var a [3]int                    // len=3, toàn zero (0)
//	a := [3]int{1, 2, 3}            // ghi rõ len
//	a := [...]int{10, 20, 30}       // KHÔNG ghi len — compiler đếm phần tử → [3]int
//	a := [5]int{1, 2}               // len=5, phần còn lại = 0
//	a := [5]int{2: 99, 4: 1}        // gán theo index; chỗ trống = 0
//
// Lưu ý: [...] chỉ dùng khi khai báo + literal cùng lúc; len vẫn cố định sau compile.
package gobasics

func ArraySum(a [5]int) int {
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}

func ArrayLiteral() [3]string {
	return [3]string{"a", "b", "c"}
}

// ArrayInferredLen — [...]int{...}: không viết số trong [ ], compiler suy ra len.
func ArrayInferredLen() [4]int {
	return [...]int{10, 20, 30, 40}
}

// ArrayZeroInit — var + zero values.
func ArrayZeroInit() [3]int {
	var a [3]int
	return a
}

// ArrayPartialInit — khai báo len lớn hơn số phần tử trong literal.
func ArrayPartialInit() [5]int {
	return [5]int{1, 2}
}

// ArrayIndexInit — gán tại index cụ thể (các ô khác = zero).
func ArrayIndexInit() [5]int {
	return [5]int{2: 99, 4: 1}
}

// ArrayCopyDemo — array là value type: gán = copy toàn bộ.
func ArrayCopyDemo() (original, copied [3]int) {
	original = [3]int{1, 2, 3}
	copied = original
	copied[0] = 99
	return original, copied
}
