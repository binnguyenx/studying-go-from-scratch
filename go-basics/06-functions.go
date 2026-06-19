// Functions — params, multiple returns, named returns, variadic, first-class
//
// func f(a, b int) (int, error)
// func f() (x int) { return }  // naked return
// func sum(xs ...int) int
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
