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

func main() {
	// TODO: Warm-up — name, age, gpa, one fmt.Printf line (import "fmt").

	// TODO: Main exercise — "Course roster" per spec above (Println + Printf + Sprintf).
}
