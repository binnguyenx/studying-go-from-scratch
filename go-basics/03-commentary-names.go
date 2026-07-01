// Effective Go — Commentary & Names
// https://go.dev/doc/effective_go#commentary
// https://go.dev/doc/effective_go#names
//
// Exported names: MixedCaps (not snake_case). Getters omit "Get" prefix.
// Package names: short, lowercase, no underscores.

package gobasics

// Example exported name (MixedCaps).
func ExportedCounter() int { return 1 }
