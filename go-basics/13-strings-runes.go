// Strings & runes — UTF-8, len bytes, range gives runes, string immutable
//
// s := "hello"
// for i, r := range s { }
// string([]byte{104,105})
package gobasics

func StringLenBytes(s string) int {
	return len(s)
}

func RuneCount(s string) int {
	n := 0
	for range s {
		n++
	}
	return n
}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}
