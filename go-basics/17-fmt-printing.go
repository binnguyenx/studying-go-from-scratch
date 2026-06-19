// fmt — Println, Printf, Sprintf (import "fmt" in real programs)
//
// Verbs: %s %d %f %.2f %t %v %T %% \n
// package gobasics does not import fmt here; see syntax_test.go for output checks.
package gobasics

// FormatLine builds a line like fmt.Sprintf("%s %d %.1f", name, age, score).
func FormatLine(name string, age int, score float64) string {
	// minimal sprintf-style without fmt import for teaching
	return name + " age=" + itoa(age) + " score=" + floatString(score, 1)
}

func floatString(f float64, decimals int) string {
	if decimals <= 0 {
		return itoa(int(f))
	}
	neg := f < 0
	if neg {
		f = -f
	}
	scale := 1
	for i := 0; i < decimals; i++ {
		scale *= 10
	}
	whole := int(f)
	frac := int((f-float64(whole))*float64(scale) + 0.5)
	s := itoa(whole) + "." + padLeft(itoa(frac), decimals, '0')
	if neg {
		return "-" + s
	}
	return s
}

func padLeft(s string, n int, c byte) string {
	for len(s) < n {
		s = string(c) + s
	}
	return s
}
