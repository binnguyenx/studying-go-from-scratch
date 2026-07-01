// Effective Go — Maps
// https://go.dev/doc/effective_go#maps
//
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


func CopyMapRange(oldMap map[string]int) map[string]int {
	newMap := make(map[string]int, len(oldMap))
	for key, value := range oldMap {
		newMap[key] = value
	}
	return newMap
}

func MapKeys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
