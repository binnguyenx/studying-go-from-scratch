// Packages & imports — one package per directory, exported = Capital letter
//
// import "fmt"
// import f "fmt"
// import ( "fmt"; "os" )
// go mod init module/name at repo root
package gobasics

// ExportedName is visible outside package gobasics.
func ExportedName() string {
	return "exported"
}

// unexportedHelper is only visible inside gobasics.
func unexportedHelper() string {
	return "internal"
}
