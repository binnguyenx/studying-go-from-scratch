// defer — runs when surrounding function returns (LIFO stack)
//
// defer f()
// defer fmt.Println("cleanup")
// common: close file, unlock mutex, recover in defer
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
