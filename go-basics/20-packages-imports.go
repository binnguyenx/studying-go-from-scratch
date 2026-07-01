// Effective Go — packages
// https://go.dev/doc/effective_go
//
package gobasics

// ExportedName is visible outside package gobasics.
func ExportedName() string {
	return "exported"
}

// unexportedHelper is only visible inside gobasics.
func unexportedHelper() string {
	return "internal"
}
