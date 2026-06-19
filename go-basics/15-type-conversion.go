// Type conversion — explicit only (no implicit numeric mixing)
//
// float64(i)   int(f)   string(b) via strconv
// cannot convert []int to []float64 element-wise without loop
package gobasics

func IntToFloat(n int) float64 {
	return float64(n)
}

func FloatToInt(f float64) int {
	return int(f) // truncates toward zero
}

func BytesToString(b []byte) string {
	return string(b)
}

func StringToBytes(s string) []byte {
	return []byte(s)
}
