// Basic types & zero values
//
// bool | string | int int8 int16 int32 int64 | uint ... | float32 float64
// | complex64 complex128 | byte (=uint8) | rune (=int32)
//
// Zero: 0, "", false, nil (pointer, slice, map, chan, func, interface)
package gobasics

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
