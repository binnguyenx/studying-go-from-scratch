// Effective Go — Initialization (constants & variables)
// https://go.dev/doc/effective_go#constants
// https://go.dev/doc/effective_go#variables
package gobasics

// Package-level
var AppName = "go-basics"

const Version = "1.0"

const (
	StatusOK = 200
	MaxRetry = 3
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
)

// DemoVars returns sample values declared with different styles.
func DemoVars() (language string, age int, ratio float64, ok bool) {
	var count int
	var msg string
	var ready bool

	language = "Go"
	age = 21
	ratio = 3.14
	ok = true
	count = 10
	msg = "hi"
	ready = false

	_ = count
	_ = msg
	_ = ready

	short := 42
	_ = short

	return language, age, ratio, ok
}

// ZeroValues demonstrates default values before assignment.
func ZeroValues() (i int, f float64, s string, b bool) {
	return // named returns stay at zero values
}

// TypeSizes returns names of common types (for learning, not sizeof).
func TypeNames() []string {
	return []string{
		"bool", "string", "int", "int64", "uint", "float64",
		"byte", "rune", "complex128",
	}
}

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

func IntToFloat(n int) float64 {
	return float64(n)
}

func FloatToInt(f float64) int {
	return int(f) // truncates toward zero
}

func BytesToString(b []byte) string {
	return string(b)
}

func StringToBytes(s string) []byte {
	return []byte(s)
}
