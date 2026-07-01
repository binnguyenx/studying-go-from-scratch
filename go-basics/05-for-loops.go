// for — the only loop in Go (Effective Go: https://go.dev/doc/effective_go#for)
//
// Go unifies for and while; there is no do-while. Three forms (only C-style uses semicolons):
//
//   for init; condition; post { }   // like C for
//   for condition { }               // like C while
//   for { }                         // like C for(;;) — use break/return to exit
//
// Short declaration in the loop header is idiomatic:
//
//   for i := 0; i < 10; i++ { sum += i }
//
// range manages loops over array, slice, string, map, or channel:
//
//   for key, value := range m { }
//   for key := range m { }            // only key / index
//   for _, value := range s { }       // discard index with _
//
// Strings: range yields index (bytes) and rune (Unicode code point), not byte index per char.
// See RangeStringRunes — invalid UTF-8 becomes U+FFFD (one byte consumed).
//
// Go has no comma operator; ++ and -- are statements, not expressions.
// Multiple variables in for init/post use parallel assignment:
//
//   for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 { ... }  // reverse in place
//
// Also read: https://go.dev/tour/flowcontrol/1  https://go.dev/ref/spec#For_statements
package gobasics

// SumRange — C-style for (Effective Go sum 0..9 pattern).
func SumRange(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum
}

// CountWhile — for condition { } (no do-while in Go).
func CountWhile(n int) int {
	count := 0
	for n > 0 {
		count++
		n /= 2
	}
	return count
}

// FirstEven — range slice; only need value → for _, v := range.
func FirstEven(nums []int) (int, bool) {
	for _, v := range nums {
		if v%2 == 0 {
			return v, true
		}
	}
	return 0, false
}

// CopyMapRange — range map (key, value); map iteration order is undefined.
func CopyMapRange(oldMap map[string]int) map[string]int {
	newMap := make(map[string]int, len(oldMap))
	for key, value := range oldMap {
		newMap[key] = value
	}
	return newMap
}

// SumRangeValues — range with _ when only values matter.
func SumRangeValues(nums []int) int {
	sum := 0
	for _, value := range nums {
		sum += value
	}
	return sum
}

// RangeStringRunes — Effective Go string range: byte index + rune (UTF-8 decoded).
type RuneAt struct {
	Pos int
	R   rune
}

func RangeStringRunes(s string) []RuneAt {
	out := make([]RuneAt, 0, len(s))
	for pos, char := range s {
		out = append(out, RuneAt{Pos: pos, R: char})
	}
	return out
}

// ReverseSliceParallel — Effective Go reverse using parallel assignment in for header.
func ReverseSliceParallel(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

// BreakFromInfinite — for { } with break (no do-while).
func BreakFromInfinite(limit int) int {
	sum := 0
	for {
		if sum >= limit {
			break
		}
		sum++
	}
	return sum
}
