// Variables & constants — var, :=, const, iota
//
// Syntax:
//   var name type = value
//   var name type          // zero value
//   name := value          // short decl (inside functions only)
//   const Pi = 3.14
//   const ( A = 1; B = iota ) // iota resets per const block
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
