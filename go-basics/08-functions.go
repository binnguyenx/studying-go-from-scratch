// Effective Go — functions.go
// https://go.dev/doc/effective_go
//
package gobasics

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Divide(a, b int) (quotient int, ok bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

func bounds(a, b int) (min, max int) {
	if a <= b {
		min, max = a, b
	} else {
		min, max = b, a
	}
	return
}

func SumVariadic(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func ApplyTwice(f func(int) int, x int) int {
	return f(f(x))
}

func Double(x int) int { return x * 2 }
