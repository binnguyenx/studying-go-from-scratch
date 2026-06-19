// Operators — arithmetic, comparison, logical, bitwise
//
// Arithmetic: + - * / % (ints only for %)
// Comparison: == != < <= > >=
// Logical: && || !   (short-circuit)
// Bitwise: & | ^ << >> &^
package gobasics

func Add(a, b int) int       { return a + b }
func Sub(a, b int) int       { return a - b }
func Mul(a, b int) int       { return a * b }
func Div(a, b int) int       { return a / b }
func Mod(a, b int) int       { return a % b }
func Equal(a, b int) bool    { return a == b }
func And(a, b bool) bool    { return a && b }
func Or(a, b bool) bool     { return a || b }
func Not(a bool) bool       { return !a }
func BitAnd(a, b int) int   { return a & b }
func ShiftLeft(a, n int) int { return a << n }
