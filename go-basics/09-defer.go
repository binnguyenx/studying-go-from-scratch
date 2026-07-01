// Effective Go — defer.go
// https://go.dev/doc/effective_go
//
package gobasics

func DeferOrder() (order []int) {
	defer func() { order = append(order, 3) }()
	defer func() { order = append(order, 2) }()
	order = append(order, 1)
	return
}

// DeferSum uses defer to accumulate after main body logic.
func DeferSum(nums []int) (sum int) {
	defer func() { sum *= 2 }() // runs last: doubles result
	for _, n := range nums {
		sum += n
	}
	return sum
}
