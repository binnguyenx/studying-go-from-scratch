// for — the only loop in Go (C-style, while-style, infinite, range)
//
// for i := 0; i < n; i++ { }
// for cond { }
// for { break }
// for i, v := range xs { }
package gobasics

func SumRange(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func CountWhile(n int) int {
	count := 0
	for n > 0 {
		count++
		n /= 2
	}
	return count
}

func FirstEven(nums []int) (int, bool) {
	for _, v := range nums {
		if v%2 == 0 {
			return v, true
		}
	}
	return 0, false
}
