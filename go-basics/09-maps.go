// Maps — map[K]V, make, delete, ok idiom, not safe for concurrent use
//
// m := map[string]int{}
// m := make(map[string]int)
// v, ok := m[key]
// delete(m, key)
package gobasics

func WordCount(s string) map[string]int {
	counts := make(map[string]int)
	for _, w := range splitWords(s) {
		counts[w]++
	}
	return counts
}

func splitWords(s string) []string {
	var out []string
	start := -1
	for i := 0; i <= len(s); i++ {
		if i < len(s) && s[i] != ' ' {
			if start < 0 {
				start = i
			}
			continue
		}
		if start >= 0 {
			out = append(out, s[start:i])
			start = -1
		}
	}
	return out
}
