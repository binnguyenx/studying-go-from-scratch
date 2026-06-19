// Arrays — fixed length [N]T, value type (copied on assign)
//
// var a [3]int
// a := [3]int{1,2,3}
// a := [...]int{1,2,3}  // compiler counts
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
