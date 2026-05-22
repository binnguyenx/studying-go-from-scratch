// Day 1 — Setup + Basics
//
// Aligns with README Day 1: install Go, VS Code + Go extension, variables,
// types (string, int, float64, bool), fmt.Println (and Printf for formatting).
//
// -----------------------------------------------------------------------------
// # 1. Install Go
// -----------------------------------------------------------------------------
// - Download: https://go.dev/dl/
// - Verify in terminal:
//     go version
// - You should see go1.21+ (or newer). No need to memorize GOPATH for daily work.
//
// -----------------------------------------------------------------------------
// # 2. Editor — VS Code + Go extension
// -----------------------------------------------------------------------------
// - Install Visual Studio Code: https://code.visualstudio.com/
// - Install extension **Go** (publisher: Go Team at Google):
//     https://marketplace.visualstudio.com/items?itemName=golang.go
// - Open this repo folder; the extension offers formatting, navigation, and `go run` hints.
//
// -----------------------------------------------------------------------------
// # 3. Variables — three ways you will use today
// -----------------------------------------------------------------------------
// | Style            | Example              | Typical use                          |
// |------------------|----------------------|--------------------------------------|
// | var              | var name string      | Package level, explicit zero value   |
// | short declaration| age := 21            | Inside func; type inferred           |
// | const            | const MaxScore = 100| Values fixed at compile time         |
//
// Zero values (no initializer): numeric 0, string "", bool false, pointers nil.
//
// -----------------------------------------------------------------------------
// # 4. Types for today (built-in)
// -----------------------------------------------------------------------------
// - string   — text (UTF-8); len(s) is bytes, not runes.
// - int      — whole numbers; also int8…int64 exist; default integer type is int.
// - float64  — floating point; default for floating literals like 3.14.
// - bool     — only true or false.
//
// Type conversion is explicit: float64(age), int(gpa) — no implicit mixing.
//
// -----------------------------------------------------------------------------
// # 5. Printing — fmt
// -----------------------------------------------------------------------------
// - fmt.Println(a, b, c) — spaces between operands, then newline.
// - fmt.Printf(format, args...) — format string uses verbs, e.g. %s %d %.2f %t %% \n
// - fmt.Sprintf — same as Printf but returns a string (useful before Println).
//
// Package doc: https://pkg.go.dev/fmt
//
// -----------------------------------------------------------------------------
// # 6. Minimal program shape
// -----------------------------------------------------------------------------
//   package main
//
//   import "fmt"
//
//   func main() {
//   	// entry point of a command-line program
//   }
//
// -----------------------------------------------------------------------------
// # 7. Read / skim (after setup)
// -----------------------------------------------------------------------------
// - Tour of Go — Packages, Variables, Functions: https://go.dev/tour/list
// - Go Proverbs: https://go-proverbs.github.io/
// - Pitfalls (skim): https://golang50shades.com/
//
// =============================================================================
// # Warm-up (small)
// =============================================================================
// Declare three variables:
//   - name  string
//   - age   int
//   - gpa   float64
// Print **one** line using a single fmt.Printf with format verbs (e.g. %s, %d, %.2f).
//
// Hint: one format string, three arguments after it.
//
// =============================================================================
// # Main exercise — "Course roster" (larger, still Day-1 only)
// =============================================================================
// Build a small **text report** in main using only variables, constants, fmt.Println,
// fmt.Printf, and fmt.Sprintf (no structs, slices, loops, or user input unless you
// already know them — the point is types + printing).
//
// **Data you must model (all used in output):**
//   - const: course code (string, e.g. "CS-100") and room capacity (int).
//   - string: course title, instructor name, building/room label.
//   - int:    enrolled count, waitlist count.
//   - float64: average pre-test score (e.g. 72.5).
//   - bool:   whether registration is still open.
//
// **Output requirements:**
//   1) A title block using fmt.Println (several lines) with blank line separation.
//   2) One paragraph line built with fmt.Sprintf stored in a variable, then printed.
//   3) A small "table" of two **fictional** students:
//        - Each row: last name (string), age (int), lab score (float64), passed (bool).
//      Use fmt.Printf with **padding** so columns line up (e.g. %-12s %3d %6.1f %5t).
//   4) A footer line showing enrollment vs capacity using int variables and %d.
//   5) One final fmt.Printf that includes **at least four** different verbs in one call
//      (e.g. %s %d %.2f %t) summarizing the course.
//
// **Quality bar:** someone reading stdout should understand the course snapshot without
// reading your source.
//
// Run from repo root:
//
//	go run ./day01
//
// =============================================================================
// # Checklist
// =============================================================================
// [ ] go version works
// [ ] VS Code + Go extension installed
// [ ] Warm-up: one Printf line with name, age, gpa
// [ ] Main exercise: all required types and print sections present
//
package main

import "fmt"

func main() {
	// --- Demo: common ways to declare variables in Go (run: go run ./day01) ---

	// 1) const — compile-time constant (typed or untyped)
	const pi = 3.14159
	const (
		StatusOK    = 200
		StatusNotFound = 404
	)
	const (
		RoleGuest = iota // 0
		RoleUser         // 1
		RoleAdmin        // 2
	)

	// 2) var — explicit type + value
	var language string = "Go"
	var year int = 2009

	// 3) var — explicit type, zero value (int → 0, string → "", bool → false)
	var count int
	var name string
	var ready bool

	// 4) var — type inferred from literal
	var ratio = 0.5 // float64

	// 5) short declaration (only inside functions)
	msg := "hello"
	score := 42

	// 6) var block — several at once
	var (
		host     = "localhost"
		port int = 8080
	)

	// 7) multiple names, one type
	var x, y, z int = 1, 2, 3

	// 8) parallel short declaration (different types ok)
	a, b, c := 1, 2.5, true

	// 9) pointer variable (value is nil until assigned)
	var ptr *int
	n := 100
	ptr = &n

	fmt.Println("=== const ===")
	fmt.Printf("pi=%v StatusOK=%d StatusNotFound=%d RoleAdmin=%d\n", pi, StatusOK, StatusNotFound, RoleAdmin)

	fmt.Println("\n=== var (explicit + zero values) ===")
	fmt.Printf("language=%q year=%d count=%d name=%q ready=%t\n", language, year, count, name, ready)

	fmt.Println("\n=== inferred + short decl ===")
	fmt.Printf("ratio=%T %v msg=%q score=%d\n", ratio, ratio, msg, score)

	fmt.Println("\n=== var block + multi var ===")
	fmt.Printf("host=%s port=%d x,y,z=%d,%d,%d\n", host, port, x, y, z)

	fmt.Println("\n=== parallel :=  + pointer ===")
	fmt.Printf("a=%d b=%.1f c=%t ptr=%v *ptr=%d\n", a, b, c, ptr, *ptr)
}
