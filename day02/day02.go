// Day 2 — Control flow + Functions
//
// Lý thuyết tóm tắt (Markdown): theory.md trong cùng thư mục.
//
// Khớp README Day 2: if, switch, for (không có while), function basics, named returns.
//
// -----------------------------------------------------------------------------
// # 1. if
// -----------------------------------------------------------------------------
// - Không bắt buộc ngoặc quanh điều kiện: if x > 0 { }
// - Khối init + điều kiện: if err := do(); err != nil { ... } — phạm vi biến chỉ trong if/else.
//
// -----------------------------------------------------------------------------
// # 2. switch
// -----------------------------------------------------------------------------
// - switch x { case a: case b: default: } — không rơi xuống case sau (trừ fallthrough).
// - switch { case cond: } — “tagless”, thay cho if-else dài.
//
// -----------------------------------------------------------------------------
// # 3. for (mọi vòng lặp)
// -----------------------------------------------------------------------------
// - C-style: for i := 0; i < n; i++ { }
// - “while”: for cond { }
// - Vô hạn: for { break }
// - range: for i, v := range items { } (slice/string/map…)
//
// -----------------------------------------------------------------------------
// # 4. Functions
// -----------------------------------------------------------------------------
// - func Name(in type) outtype { }
// - Named returns: func f() (x, y int) { x=1; y=2; return }
// - Nhiều tham số cùng kiểu: func add(a, b int) int
//
// -----------------------------------------------------------------------------
// # 5. Read
// -----------------------------------------------------------------------------
// - Tour: https://go.dev/tour/list — For, If, Switch, Functions
//
// =============================================================================
// # Warm-up
// =============================================================================
// Viết Max(a, b int) int trả về số lớn hơn (chỉ dùng if).
//
// =============================================================================
// # Bài chính (README)
// =============================================================================
// - FizzBuzz(n int) string: chia hết cho 15 → "FizzBuzz", 3 → "Fizz", 5 → "Buzz", else strconv.Itoa(n).
// - Trong main: in FizzBuzz cho từng số 1..30 (một dòng mỗi số).
//
// =============================================================================
// # Bài lớn — “Báo cáo số”
// =============================================================================
// 1) PrintFizzBuzzRange(lo, hi int): vòng for từ lo đến hi, mỗi dòng: n + tab + FizzBuzz(n).
// 2) Sign(n int) int: trả -1 / 0 / 1 (if hoặc switch).
// 3) bounds(a, b int) (min, max int) — named return: min max đúng thứ tự nhỏ → lớn.
// 4) classifyHour(h int) string: switch — 0–11 "AM block", 12–17 "afternoon", 18–21 "evening", else "night"
//    (xử lý giờ ngoài 0–23 tùy bạn: clamp hoặc "invalid").
// 5) main: gọi đủ hàm trên, in rõ từng mục (Println tiêu đề + kết quả).
//
// Chạy: go run ./day02
//
package main

import (
	"fmt"
	"strconv"
)

// Max returns the larger of a and b.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// FizzBuzz classic mapping for one integer.
func FizzBuzz(n int) string {
	switch {
	case n%15 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(n)
	}
}

// PrintFizzBuzzRange prints n and FizzBuzz(n) for each n in [lo, hi].
func PrintFizzBuzzRange(lo, hi int) {
	for n := lo; n <= hi; n++ {
		fmt.Printf("%d\t%s\n", n, FizzBuzz(n))
	}
}

// Sign returns -1, 0, or 1.
func Sign(n int) int {
	switch {
	case n < 0:
		return -1
	case n == 0:
		return 0
	default:
		return 1
	}
}

// bounds returns min and max using named results.
func bounds(a, b int) (min, max int) {
	if a <= b {
		min, max = a, b
	} else {
		min, max = b, a
	}
	return
}

func classifyHour(h int) string {
	switch {
	case h < 0 || h > 23:
		return "invalid"
	case h <= 11:
		return "AM block"
	case h <= 17:
		return "afternoon"
	case h <= 21:
		return "evening"
	default:
		return "night"
	}
}

func main() {
	fmt.Println("=== Warm-up: Max ===")
	fmt.Println(Max(3, 7), Max(10, 2))

	fmt.Println("\n=== README: 1..30 FizzBuzz ===")
	for n := 1; n <= 30; n++ {
		fmt.Println(FizzBuzz(n))
	}

	fmt.Println("\n=== Large: range 1..15 tab + FizzBuzz ===")
	PrintFizzBuzzRange(1, 15)

	fmt.Println("\n=== Sign samples ===")
	for _, v := range []int{-8, 0, 42} {
		fmt.Printf("Sign(%d) = %d\n", v, Sign(v))
	}

	fmt.Println("\n=== Named returns: bounds(8, 3) ===")
	fmt.Println(bounds(8, 3))

	fmt.Println("\n=== classifyHour (switch) ===")
	for _, h := range []int{9, 14, 19, 23, 99} {
		fmt.Printf("%02d -> %s\n", h, classifyHour(h))
	}
}
